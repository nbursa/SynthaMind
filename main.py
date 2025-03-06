import time
from world_simulation.wikipedia_explorer import WikipediaExplorer

# Initialize Wikipedia knowledge system
wiki_explorer = WikipediaExplorer()

print("EvolvAI is starting...")

# Ask user if EvolvAI should explore on its own
auto_exploration = input("Enable auto-exploration? (yes/no): ").strip().lower() == "yes"

while True:
    if not auto_exploration:
        topic = input("Enter a topic for EvolvAI to study (or press Enter to stop): ").strip()
        if not topic:
            print("🚀 EvolvAI has completed its study.")
            break
    else:
        topic = wiki_explorer.suggest_next_topic() or input("Enter a new topic to guide EvolvAI: ").strip()
        if not topic:
            print("🚀 EvolvAI has exhausted available knowledge.")
            break

    summary = wiki_explorer.fetch_summary(topic)

    if summary:
        print(f"📖 EvolvAI learned: {summary[:300]}...")
        related_topics = wiki_explorer.explore_related_topics(topic)

        if related_topics:
            print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

    time.sleep(1)
