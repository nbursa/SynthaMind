import time
from world_simulation.wikipedia_explorer import WikipediaExplorer
from core_intelligence.knowledge_base import KnowledgeBase

def main():
    print("EvolvAI is starting...")
    specialization = input("Enter a specialization field for EvolvAI (e.g., Medicine, Physics, AI): ").strip()
    if not specialization:
        print("Specialization is required to proceed.")
        return

    enable_auto_exploration = input("Enable auto-exploration? (yes/no): ").strip().lower() == 'yes'
    
    wiki_explorer = WikipediaExplorer()
    knowledge_base = KnowledgeBase()

    while True:
        if enable_auto_exploration:
            topic = wiki_explorer.suggest_next_topic(specialization, knowledge_base)
            if not topic:
                print("No more topics to explore within the specialization.")
                break
            print(f"🌍 EvolvAI is studying: {topic}")
        else:
            topic = input("Enter a new topic to guide EvolvAI (or press Enter to stop): ").strip()
            if not topic:
                print("🚀 EvolvAI has completed its study.")
                break

        summary = wiki_explorer.fetch_summary(topic)
        if summary:
            knowledge_base.add_entry(topic, summary)
            print(f"📖 EvolvAI learned: {summary[:500]}...")  # Displaying first 500 characters
        else:
            print(f"No information found for topic: {topic}")

        related_topics = wiki_explorer.explore_related_topics(topic, specialization)
        if related_topics:
            print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

        if enable_auto_exploration:
            user_input = input("\n📌 EvolvAI has reached a learning checkpoint.\n"
                               "Continue auto-learning (yes), select a field (no), or stop (exit)? ").strip().lower()
            if user_input == 'no':
                enable_auto_exploration = False
            elif user_input == 'exit':
                print("🚀 EvolvAI has completed its study.")
                break
        else:
            user_input = input("Do you want to continue exploring related topics? (yes/no): ").strip().lower()
            if user_input != 'yes':
                print("🚀 EvolvAI has completed its study.")
                break

        time.sleep(1)

if __name__ == "__main__":
    main()
