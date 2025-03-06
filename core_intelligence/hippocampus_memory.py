import json
import os

class HippocampusMemory:
    """Manages learned knowledge storage and retrieval."""

    def __init__(self, storage_file="knowledge_base.json"):
        self.storage_file = storage_file
        self.knowledge = self._load_memory()

    def _load_memory(self):
        """Loads memory from file or initializes an empty structure."""
        if os.path.exists(self.storage_file):
            with open(self.storage_file, "r") as file:
                return json.load(file)
        return {}

    def store_knowledge(self, topic, summary):
        """Stores learned knowledge in memory."""
        self.knowledge[topic.lower()] = summary
        self._save_memory()

    def has_learned(self, topic):
        """Checks if a topic is already learned."""
        return topic.lower() in self.knowledge

    def get_learned_topics(self):
        """Returns all learned topics."""
        return list(self.knowledge.keys())  # ✅ Fix: Ensure we return a list of learned topics

    def _save_memory(self):
        """Saves memory to file."""
        with open(self.storage_file, "w") as file:
            json.dump(self.knowledge, file, indent=4)
