import wikipediaapi

class WikipediaExplorer:
    """
    EvolvAI connects to Wikipedia, learns, and ensures it stays within a chosen field.
    """

    def __init__(self, language="en", checkpoint_interval=5):
        self.wiki = wikipediaapi.Wikipedia(
            language=language,
            user_agent="EvolvAI/1.0 (https://github.com/nbursa/evolvai; contact: nenad@nenadbursac.com)"
        )
        self.learned_topics = {}  # Stores learned topics
        self.specialization_field = None  # User-selected specialization
        self.checkpoint_interval = checkpoint_interval  # When to pause and ask for direction
        self.topics_learned_count = 0  # Counter for breakpoints

    def set_specialization(self, field):
        """ Define the specialization field EvolvAI should focus on. """
        self.specialization_field = field.lower()
        print(f"🎯 EvolvAI will focus on: {self.specialization_field.capitalize()}")

    def fetch_summary(self, topic):
        """ Fetch and store Wikipedia knowledge while ensuring specialization. """
        if topic in self.learned_topics:
            print(f"✅ EvolvAI already knows about '{topic}'.")
            return None  # Skip if already learned

        page = self.wiki.page(topic)

        if not page.exists():
            print(f"❌ Wikipedia page for '{topic}' not found.")
            return None

        summary = page.summary[:1000]  # Limit summary length
        self.learned_topics[topic] = summary  # Store learned topic
        self.topics_learned_count += 1  # Track how many topics were learned
        print(f"🌍 EvolvAI is studying: {page.title}")

        return summary

    def explore_related_topics(self, topic):
        """ Retrieve **filtered** related topics from Wikipedia to keep learning focused. """
        page = self.wiki.page(topic)
        if not page.exists():
            return []

        related = [t for t in list(page.links.keys()) if t not in self.learned_topics]

        if self.specialization_field:
            # Keep only topics that match the specialization field
            related = [t for t in related if self.specialization_field in t.lower()]

        return related[:5]  # Return first 5 relevant topics

    def should_prompt_for_direction(self):
        """ Returns True if a learning checkpoint is reached. """
        return self.topics_learned_count % self.checkpoint_interval == 0 and self.topics_learned_count > 0
    
    def suggest_next_topic(self):
        """ Suggest the next topic based on learned knowledge within the specialization field. """
        if not self.learned_topics:
            return None  # No previous learning

        last_learned = list(self.learned_topics.keys())[-1]
        related = self.explore_related_topics(last_learned)

        # ✅ Ensure the topic is within specialization
        if self.specialization_field:
            related = [t for t in related if self.specialization_field in t.lower()]

        return related[0] if related else None  # Pick first relevant topic
