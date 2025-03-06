import wikipediaapi

class WikipediaExplorer:
    """Handles Wikipedia-based knowledge retrieval and topic exploration."""

    def __init__(self, language='en', user_agent="EvolvAI/1.0 (contact: your-email@example.com)"):
        """Initialize Wikipedia API with the required user-agent."""
        self.wiki = wikipediaapi.Wikipedia(user_agent=user_agent, language=language)

    def fetch_summary(self, topic):
        """Retrieve a summary from Wikipedia."""
        page = self.wiki.page(topic)
        if page.exists():
            return page.summary
        return None

    def explore_related_topics(self, topic, specialization):
        """Suggest related topics within the chosen specialization."""
        page = self.wiki.page(topic)
        if not page.exists():
            return []

        related_topics = [link_title for link_title in page.links.keys()
                          if specialization.lower() in link_title.lower()]

        return related_topics[:5]  # Return top 5 relevant topics

    def suggest_next_topic(self, specialization, knowledge_base):
        """Suggest the next topic based on the specialization."""
        specialization_page = self.wiki.page(specialization)
        if not specialization_page.exists():
            return None

        for link_title in specialization_page.links.keys():
            if specialization.lower() in link_title.lower() and not knowledge_base.has_learned(link_title):
                return link_title
        return None
