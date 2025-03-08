"""
🚀 SynthaMind - ChromaDB Initialization
This script starts a persistent ChromaDB server for vector storage.
"""

import chromadb

# Initialize and run ChromaDB with persistent storage
chroma_client = chromadb.PersistentClient(path="./chromadb_data")

print("✅ ChromaDB server is running with persistent storage at './chromadb_data'")
