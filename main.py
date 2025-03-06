import time
from world_simulation.wikipedia_explorer import WikipediaExplorer
from core_intelligence.knowledge_base import KnowledgeBase

# Initialize the knowledge base
knowledge_base = KnowledgeBase()

# Initialize Wikipedia knowledge system
wiki_explorer = WikipediaExplorer()

def main():
    print("EvolvAI is starting...")
    
    # Get specialization field from the user
    specialization = input("Enter a specialization field for EvolvAI (e.g., Medicine, Physics, AI): ").strip()
    knowledge_base.add_entry(specialization, "")  # Ensure specialization is known
    
    # Enable auto-exploration
    auto_explore = input("Enable auto-exploration? (yes/no): ").strip().lower() == 'yes'

    while True:
        if auto_explore:
            next_topic = wiki_explorer.suggest_next_topic(specialization, knowledge_base)
            if not next_topic:
                print("No more topics to explore.")
                break

            print(f"🌍 EvolvAI is studying: {next_topic}")
            summary = wiki_explorer.fetch_summary(next_topic)
            if summary:
                print(f"📖 EvolvAI learned: {summary[:300]}...")  # Show a preview
                knowledge_base.add_entry(next_topic, summary)

            # Suggest related topics for further exploration
            related_topics = wiki_explorer.explore_related_topics(next_topic, specialization)
            if related_topics:
                print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

            # Ask if the user wants to continue auto-exploring
            continue_auto = input("Continue auto-learning (yes), select a field (no), or stop (exit)? ").strip().lower()
            if continue_auto == "exit":
                break
            elif continue_auto == "no":
                specialization = input("Enter a new specialization or a topic: ").strip()
                knowledge_base.add_entry(specialization, "")  # Add a new field
            elif continue_auto == "yes":
                continue  # Continue the loop with auto-exploration

        else:
            topic = input("Enter a topic for EvolvAI to study (or press Enter to stop): ").strip()
            if not topic:
                print("🚀 EvolvAI has completed its study.")
                break

            summary = wiki_explorer.fetch_summary(topic)
            if summary:
                print(f"📖 EvolvAI learned: {summary[:300]}...")  # Show a preview of knowledge
                knowledge_base.add_entry(topic, summary)

            # Suggest related topics for further exploration
            related_topics = wiki_explorer.explore_related_topics(topic, specialization)
            if related_topics:
                print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

            # Ask if the user wants to continue
            continue_learning = input("Continue learning? (yes/no): ").strip().lower()
            if continue_learning != "yes":
                break

if __name__ == "__main__":
    main()
