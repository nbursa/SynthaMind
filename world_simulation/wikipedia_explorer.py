import wikipediaapi
import requests

class WikipediaExplorer:
    """Handles Wikipedia-based knowledge retrieval and topic exploration."""

    def __init__(self, language='en', user_agent='EvolvAI/1.0'):
        # Directly specify user_agent in the constructor
        self.wiki = wikipediaapi.Wikipedia(
            language=language, 
            user_agent=user_agent  # Specify user_agent directly
        )

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

        related_topics = []
        for link_title in page.links.keys():
            if specialization.lower() in link_title.lower():
                related_topics.append(link_title)

        return related_topics[:5]  # Return top 5 relevant topics

    def suggest_next_topic(self, specialization, knowledge_base):
        """Suggest the next topic based on the specialization."""
        for page in self.wiki.page(specialization).links.keys():
            if specialization.lower() in page.lower() and not knowledge_base.has_learned(page):
                return page
        return None
