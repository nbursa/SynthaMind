from world_simulation.wikipedia_explorer import WikipediaExplorer
from core_intelligence.hippocampus_memory import HippocampusMemory
import random

class ThalamusController:
    """Routes knowledge between instinct processing, memory, and specialization engine."""

    def __init__(self, specialization="general"):
        self.memory = HippocampusMemory()
        self.wiki_explorer = WikipediaExplorer()
        self.specialization = specialization.lower()
        self.interest_levels = {}  # Tracks interest intensity in topics
        self.instinct_priority = ["machine learning", "neural networks", "deep learning"]  # Default AI instincts

    def learn_topic(self, topic):
        """AI learns a topic, checking memory first."""
        if not isinstance(topic, str) or not topic.strip():
            print("⚠️ Invalid topic! Try again.")
            return None

        topic = topic.lower()

        if self.memory.has_learned(topic):
            print(f"✅ EvolvAI already knows about {topic}.")
            return None

        summary = self.wiki_explorer.fetch_summary(topic)
        if summary:
            self.memory.store_knowledge(topic, summary)
            self.track_interest(topic)
            print(f"📖 EvolvAI learned: {summary[:300]}...")
            return summary
        else:
            print(f"⚠️ No information found on {topic}. AI will self-correct.")
            return self.suggest_next_topic()

    def suggest_next_topic(self):
        """Suggests the next best topic, ensuring it's relevant."""
        learned_topics = self.memory.get_learned_topics()

        # Prioritize top instinct-related topics
        for instinct in self.instinct_priority:
            if instinct not in learned_topics:
                return instinct

        # Find next topic aligned with specialization
        related_topics = []
        for topic in learned_topics:
            related = self.wiki_explorer.explore_related_topics(topic, self.specialization)
            related_topics.extend([t for t in related if t not in learned_topics])

        if related_topics:
            return random.choice(related_topics)  # Pick one at random to keep exploration dynamic

        print("🛑 No more topics available. EvolvAI has reached knowledge saturation.")
        return None

    def track_interest(self, topic):
        """Adjusts AI interest in topics, reinforcing learning depth."""
        self.interest_levels[topic] = self.interest_levels.get(topic, 0) + 1
