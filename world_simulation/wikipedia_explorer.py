import wikipediaapi

class WikipediaExplorer:
    """
    EvolvAI connects to Wikipedia to dynamically expand knowledge and store what it learns.
    """

    def __init__(self, language="en"):
        self.wiki = wikipediaapi.Wikipedia(
            language=language,
            user_agent="EvolvAI/1.0 (https://github.com/nbursa/evolvai; contact: youremail@example.com)"
        )
        self.learned_topics = {}  # Stores learned knowledge

    def fetch_summary(self, topic):
        """ Fetch and store Wikipedia knowledge. """
        if topic in self.learned_topics:
            print(f"✅ EvolvAI already knows about '{topic}'.")
            return None  # Skip if already learned

        page = self.wiki.page(topic)

        if not page.exists():
            print(f"❌ Wikipedia page for '{topic}' not found.")
            return None

        summary = page.summary[:1000]  # Limit summary length
        self.learned_topics[topic] = summary  # Store learned topic
        print(f"🌍 EvolvAI is studying: {page.title}")
        return summary

    def explore_related_topics(self, topic):
        """ Retrieve related topics from Wikipedia. """
        page = self.wiki.page(topic)
        if not page.exists():
            return []

        related = [t for t in list(page.links.keys()) if t not in self.learned_topics]
        return related[:5]  # Return first 5 unexplored topics

    def suggest_next_topic(self):
        """ Suggest the next topic based on learned knowledge. """
        if not self.learned_topics:
            return None  # No previous learning

        last_learned = list(self.learned_topics.keys())[-1]
        related = self.explore_related_topics(last_learned)
        return related[0] if related else None
