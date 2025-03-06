import wikipediaapi

class WikipediaExplorer:
    """
    EvolvAI connects to Wikipedia to dynamically expand knowledge.
    """

    def __init__(self, language="en"):
        self.wiki = wikipediaapi.Wikipedia(
            language=language,
            user_agent="EvolvAI/1.0 (https://github.com/nbursa/evolvai; contact: youremail@example.com)"
        )

    def fetch_summary(self, topic):
        """ Fetch the introduction summary of a Wikipedia article. """
        page = self.wiki.page(topic)

        if not page.exists():
            print(f"❌ Wikipedia page for '{topic}' not found.")
            return None

        print(f"🌍 EvolvAI is studying: {page.title}")
        return page.summary[:1000]  # Limit to 1000 characters for quick processing

    def explore_related_topics(self, topic):
        """ Retrieve related links from a Wikipedia article. """
        page = self.wiki.page(topic)

        if not page.exists():
            return []

        return list(page.links.keys())[:5]  # Return first 5 related topics
