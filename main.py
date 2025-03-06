import time
import random
from world_simulation.environment import AIWorld
from world_simulation.wikipedia_explorer import WikipediaExplorer

# Initialize both worlds
simulated_world = AIWorld(width=10, height=10)  # Simulated learning environment
wiki_explorer = WikipediaExplorer()  # Wikipedia-based knowledge system

print("EvolvAI is starting...")
print("Choose a mode:")
print("1️⃣ Simulated World Exploration")
print("2️⃣ Wikipedia Knowledge Learning")

mode = input("Enter 1 or 2: ").strip()

if mode == "1":
    print("🌍 EvolvAI will explore the simulated world.")
    while True:
        if simulated_world.has_fully_explored():
            print("🚀 EvolvAI has fully explored the simulated world. Switching to Wikipedia.")
            mode = "2"  # Auto-switch to Wikipedia when the simulated world is exhausted
            break

        valid_moves = simulated_world.get_valid_moves()
        if not valid_moves:
            print("❌ EvolvAI has no valid moves left. Exploration complete.")
            break

        random_move = random.choice(valid_moves)
        new_stimulus = simulated_world.execute_move(random_move)

        if new_stimulus:
            print(f"✅ EvolvAI found something related to {new_stimulus}!")
        else:
            print("🔄 EvolvAI is bored! Exploring the world...")

        time.sleep(1)

if mode == "2":
    print("📖 EvolvAI will now explore Wikipedia knowledge.")
    while True:
        topic = input("Enter a topic for EvolvAI to study (or press Enter to stop): ").strip()

        if not topic:
            print("🚀 EvolvAI has completed its study.")
            break

        summary = wiki_explorer.fetch_summary(topic)

        if summary:
            print(f"📖 EvolvAI learned: {summary[:300]}...")
            related_topics = wiki_explorer.explore_related_topics(topic)

            if related_topics:
                print(f"🔍 EvolvAI suggests exploring: {', '.join(related_topics)}")

        time.sleep(1)
