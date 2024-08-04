from langchain.chains import RetrievalQA
from langchain_openai import ChatOpenAI
from sentence_transformers import SentenceTransformer
from pinecone import Pinecone
from langchain.embeddings.huggingface import HuggingFaceEmbeddings
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain.document_loaders.pdf import PyPDFLoader
from langchain.vectorstores.pinecone import Pinecone as pc
from dotenv import load_dotenv, find_dotenv
import os
import openai
import sys
sys.path.append('../..')
load_dotenv(find_dotenv())
os.environ["HUGGINGFACEHUB_API_TOKEN"] = "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"
PINECONE_API_KEY = os.environ.get('PINECONE_API_KEY')
PINECONE_API_ENV = os.environ.get('ENVIRONMENT')

openai.api_key = os.environ['OPENAI_API_KEY']

loader = PyPDFLoader("./Resume-SuhailKhan.pdf")
all_docs = loader.load()

text_splitter = RecursiveCharacterTextSplitter(
    chunk_size=1500,
    chunk_overlap=150
)

splits = text_splitter.split_documents(all_docs)

model = SentenceTransformer(
    'sentence-transformers/all-mpnet-base-v2', device='cpu')
embeddings = HuggingFaceEmbeddings(
    model_name='sentence-transformers/all-mpnet-base-v2')

index_name = "pdf-chat"
pc1 = Pinecone(api_key=PINECONE_API_KEY)
index = pc1.Index(index_name)
docsearch = pc.from_texts(
    [t.page_content for t in splits], embeddings, index_name=index_name)
question = "List all the certifications"
# xq = model.encode(question).tolist()
# docs = index.query(vector=xq, top_k=3, include_metadata=True)
docs = docsearch.similarity_search(question, k=1)
print(docs[0].page_content)
llm = ChatOpenAI(model_name='gpt-3.5-turbo', temperature=0)
qa_chain = RetrievalQA.from_chain_type(
    llm=llm, chain_type="stuff", retriever=docsearch.as_retriever(search_type="similarity", search_kwargs={"k": 2}))
result = qa_chain({"query": question})
print(result["result"])
