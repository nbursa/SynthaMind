import json
import os

class KnowledgeBase:
    """ Stores learned topics and prevents redundant learning. """

    def __init__(self, filename="knowledge.json"):
        self.filename = filename
        self.learned_topics = self.load_knowledge()

    def load_knowledge(self):
        """ Load existing knowledge from a file. """
        if os.path.exists(self.filename):
            with open(self.filename, "r") as file:
                return json.load(file)
        return {}

    def save_knowledge(self):
        """ Save knowledge to a file. """
        with open(self.filename, "w") as file:
            json.dump(self.learned_topics, file, indent=4)

    def has_learned(self, topic):
        """ Check if a topic has already been learned. """
        return topic in self.learned_topics

    def add_entry(self, topic, summary):
        """ Add a new topic to the knowledge base. """
        self.learned_topics[topic] = summary
        self.save_knowledge()

    def review_learned_topics(self):
        """ Display all learned topics. """
        if not self.learned_topics:
            print("📭 EvolvAI has not learned anything yet.")
            return

        print("\n📚 EvolvAI's Learned Knowledge:")
        for idx, (topic, summary) in enumerate(self.learned_topics.items(), start=1):
            print(f"{idx}. {topic}: {summary[:200]}...")  # Show first 200 characters
