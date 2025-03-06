import json

class KnowledgeBase:
    """Stores and manages the knowledge acquired by EvolvAI."""

    def __init__(self, storage_file="knowledge_base.json"):
        self.storage_file = storage_file
        self.entries = self.load_knowledge()  # Load previously saved knowledge if available

    def add_entry(self, topic, summary):
        """Add a new topic and its summary to the knowledge base."""
        self.entries[topic] = summary
        self.save_knowledge()  # Save to file whenever new knowledge is added

    def has_learned(self, topic):
        """Check if a topic has already been learned."""
        return topic in self.entries

    def get_summary(self, topic):
        """Retrieve the summary of a learned topic."""
        return self.entries.get(topic, None)

    def save_knowledge(self):
        """Save the knowledge base to a file (in JSON format)."""
        with open(self.storage_file, 'w') as file:
            json.dump(self.entries, file, indent=4)

    def load_knowledge(self):
        """Load the knowledge base from a file."""
        try:
            with open(self.storage_file, 'r') as file:
                return json.load(file)
        except FileNotFoundError:
            return {}  # Return an empty dictionary if the file doesn't exist yet
