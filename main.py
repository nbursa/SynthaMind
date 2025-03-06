import time
from world_simulation.wikipedia_explorer import WikipediaExplorer

# Initialize Wikipedia knowledge system
wiki_explorer = WikipediaExplorer(user_agent='EvolvAI/1.0')

print("EvolvAI is starting...")

# Example specialization
specialization = 'Medicine'

while True:
    topic = input("Enter a topic for EvolvAI to study (or press Enter to stop): ").strip()
    
    if not topic:
        print("🚀 EvolvAI has completed its study.")
        break

    summary = wiki_explorer.fetch_summary(topic)
    
    if summary:
        print(f"📖 EvolvAI learned: {summary[:300]}...")  # Show preview of knowledge
        related_topics = wiki_explorer.explore_related_topics(topic, specialization)
        
        if related_topics:
            print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")
    
    time.sleep(1)
