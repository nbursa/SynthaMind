import random
from core_intelligence.hippocampus_memory import HippocampusMemory
from world_simulation.wikipedia_explorer import WikipediaExplorer

class ThalamusController:
    """Central controller that manages AI's exploration and learning decisions."""

    def __init__(self, specialization):
        self.specialization = specialization
        self.memory = HippocampusMemory()
        self.wiki_explorer = WikipediaExplorer()
        self.instinct_priority = ["machine learning", "neural networks", "deep learning"]  # Example instincts

    def learn_topic(self, topic):
        """Processes topic learning and stores knowledge."""
        if not topic or topic.lower() in ["any topic", "any ai topic"]:
            return

        summary = self.wiki_explorer.fetch_summary(topic)
        if summary:
            self.memory.store_knowledge(topic, summary)
            print(f"🌍 EvolvAI is studying: {topic}")
            print(f"📖 EvolvAI learned: {summary[:500]}...\n")  # Print preview
            related_topics = self.wiki_explorer.explore_related_topics(topic, self.specialization, self.memory.get_learned_topics())

            if related_topics:
                print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

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
            related = self.wiki_explorer.explore_related_topics(topic, self.specialization, learned_topics)
            related_topics.extend([t for t in related if t not in learned_topics])

        if related_topics:
            return random.choice(related_topics)  # Pick one at random to keep exploration dynamic

        print("🛑 No more topics available. EvolvAI has reached knowledge saturation.")
        return None
