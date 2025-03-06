import json
import os

class HippocampusMemory:
    """Handles storing and retrieving learned knowledge."""

    def __init__(self, file_path="knowledge_base.json"):
        self.file_path = file_path
        self.knowledge_base = self._load_knowledge()

    def _load_knowledge(self):
        """Loads existing knowledge from a file."""
        if os.path.exists(self.file_path):
            with open(self.file_path, "r", encoding="utf-8") as file:
                return json.load(file)
        return {}

    def store_knowledge(self, topic, summary):
        """Stores learned knowledge."""
        self.knowledge_base[topic] = summary
        self._save_knowledge()

    def get_learned_topics(self):
        """Returns topics that have been learned."""
        return list(self.knowledge_base.keys())

    def has_learned(self, topic):
        """Checks if a topic has already been learned."""
        return topic in self.knowledge_base

    def _save_knowledge(self):
        """Saves knowledge to a file."""
        with open(self.file_path, "w", encoding="utf-8") as file:
            json.dump(self.knowledge_base, file, indent=4, ensure_ascii=False)
