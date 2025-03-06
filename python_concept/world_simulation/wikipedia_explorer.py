import wikipediaapi

class WikipediaExplorer:
    """Handles Wikipedia-based knowledge retrieval and topic exploration."""

    def __init__(self, language='en', user_agent='EvolvAI/1.0'):
        self.wiki = wikipediaapi.Wikipedia(user_agent=user_agent, language=language)

    def fetch_summary(self, topic):
        """Retrieve a summary from Wikipedia."""
        page = self.wiki.page(topic)
        if page.exists():
            return page.summary
        return None

    def explore_related_topics(self, topic, specialization, learned_topics):
        """Suggest relevant related topics within the chosen specialization."""
        page = self.wiki.page(topic)
        if not page.exists():
            return []

        related_topics = []
        for link_title in page.links.keys():
            if specialization.lower() in link_title.lower() and link_title not in learned_topics:
                related_topics.append(link_title)

        return related_topics[:5]  # Return top 5 relevant topics
