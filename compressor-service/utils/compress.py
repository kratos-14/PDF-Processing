from connections.mongoConnections import MongoConnection
from gridfs import GridFSBucket
import logging
from bson import ObjectId
from io import BytesIO
import shutil
import subprocess
import os


def get_ghostscript_path():
    gs_names = ["gs", "gswin32", "gswin64"]
    for name in gs_names:
        path = shutil.which(name)
        if path is not None:
            return path
    raise FileNotFoundError(
        f"No GhostScript executable was found on path ({'/'.join(gs_names)})"
    )


def compress_pdf(name, file_ID, power=0):
    # Open the PDF file
    quality = {
        0: "/default",
        1: "/prepress",
        2: "/printer",
        3: "/ebook",
        4: "/screen"
    }
    output_file_name = name[:-4]+"_compressed.pdf"
    mongo_connection = MongoConnection()
    mongo_client = mongo_connection.get_client()
    db = mongo_client.myFiles
    fs = GridFSBucket(db)

    try:
        input_file = fs.open_download_stream(ObjectId(file_ID)).read()
        input_file_path = '/tmp/input.pdf'
        with open(input_file_path, 'wb') as f:
            f.write(input_file)

        if power < 0 or power > len(quality) - 1:
            raise Exception(f"Invalid compression level {power}")
        gs = get_ghostscript_path()
        subprocess.call(
            [
                gs,
                "-sDEVICE=pdfwrite",
                "-dCompatibilityLevel=1.4",
                "-dPDFSETTINGS={}".format(quality[power]),
                "-dNOPAUSE",
                "-dQUIET",
                "-dBATCH",
                f"-sOutputFile=/tmp/{output_file_name}",
                input_file_path,
            ]
        )
        with open(f"/tmp/{output_file_name}", 'rb') as f:
            output_file = f.read()
        fs.delete(ObjectId(file_ID))
        fs.upload_from_stream_with_id(
            ObjectId(file_ID), f"{output_file_name}", source=BytesIO(output_file))
        logging.basicConfig(level=logging.INFO)
        logging.info(f"File compressed successfully {output_file}")
        os.remove(input_file_path)
        os.remove(f"/tmp/{output_file_name}")
    except Exception as e:
        logging.basicConfig(level=logging.ERROR)
        logging.error(f"Error during compression logic: {e}")
