class KnowledgeBase:
    """Stores and manages the knowledge acquired by EvolvAI."""

    def __init__(self):
        self.entries = {}

    def add_entry(self, topic, summary):
        """Add a new topic and its summary to the knowledge base."""
        self.entries[topic] = summary

    def has_learned(self, topic):
        """Check if a topic has already been learned."""
        return topic in self.entries

    def get_summary(self, topic):
        """Retrieve the summary of a learned topic."""
        return self.entries.get(topic, None)
