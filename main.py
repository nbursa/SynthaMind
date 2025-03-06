import time
from core_intelligence.knowledge_base import KnowledgeBase
from world_simulation.wikipedia_explorer import WikipediaExplorer

# Initialize components
knowledge_base = KnowledgeBase()
wiki_explorer = WikipediaExplorer()

def main():
    print("EvolvAI is starting...")
    
    # Set specialization
    specialization = input("Enter a specialization field for EvolvAI (e.g., Medicine, Physics, AI): ").strip().lower()
    print(f"🎯 EvolvAI will focus on: {specialization.capitalize()}")

    # Enable auto-learning?
    auto_explore = input("Enable auto-exploration? (yes/no): ").strip().lower() == 'yes'

    while True:
        if auto_explore:
            topic = wiki_explorer.suggest_next_topic(specialization, knowledge_base)
        else:
            topic = input("Enter a new topic to guide EvolvAI (or press Enter to stop): ").strip().lower()
            if not topic:
                print("🚀 EvolvAI has completed its study.")
                break

        if knowledge_base.has_learned(topic):
            print(f"🔍 EvolvAI already knows about {topic.capitalize()}. Skipping...")
            continue

        print(f"🌍 EvolvAI is studying: {topic.capitalize()}")
        summary = wiki_explorer.fetch_summary(topic)
        if summary:
            print(f"📖 EvolvAI learned: {summary[:300]}...")
            knowledge_base.add_entry(topic, summary)

            # Suggest specialized related topics
            related_topics = wiki_explorer.explore_related_topics(topic, specialization)
            if related_topics:
                print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")
        else:
            print(f"❌ No information found for {topic.capitalize()}.")

        if not auto_explore:
            continue_prompt = input("\n📌 EvolvAI has reached a learning checkpoint.\n"
                                    "Continue auto-learning (yes), select a field (no), or stop (exit)? ").strip().lower()
            if continue_prompt == 'exit':
                break
            elif continue_prompt == 'no':
                auto_explore = False
            else:
                auto_explore = True

        time.sleep(1)

if __name__ == "__main__":
    main()
