import time
from core_intelligence.thalamus_controller import ThalamusController

def main():
    print("EvolvAI is starting...")

    specialization = input("Enter a specialization field for EvolvAI (e.g., Medicine, Physics, AI): ").strip().lower()
    brain = ThalamusController(specialization)

    auto_explore = input("Enable auto-exploration? (yes/no): ").strip().lower() == "yes"

    while True:
        if auto_explore:
            topic = brain.suggest_next_topic()
            if not topic:
                print("🚀 EvolvAI has fully explored its knowledge base.")
                break
        else:
            topic = input("Enter a topic to guide EvolvAI (or press Enter to stop): ").strip()

        if not topic:
            print("🚀 EvolvAI has completed its study.")
            break

        print(f"🌍 EvolvAI is studying: {topic}")
        brain.learn_topic(topic)

        print("\n📌 EvolvAI has reached a learning checkpoint.")
        choice = input("Continue auto-learning (yes), select a field (no), or stop (exit)? ").strip().lower()

        if choice == "exit":
            break
        elif choice == "no":
            auto_explore = False

        time.sleep(1)

if __name__ == "__main__":
    main()
