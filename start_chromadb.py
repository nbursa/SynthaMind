import chromadb

# Start ChromaDB server
chroma_client = chromadb.PersistentClient(path="./chromadb_data")
print("🚀 ChromaDB is running...")