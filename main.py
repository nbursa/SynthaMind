import time
from core_intelligence.thalamus_controller import ThalamusController

def main():
    print("EvolvAI is starting...")

    specialization = input("Enter a specialization field for EvolvAI (e.g., Medicine, Physics, AI): ").strip().lower()
    auto_exploration = input("Enable auto-exploration? (yes/no): ").strip().lower() == "yes"

    brain = ThalamusController(specialization)

    while True:
        topic = brain.suggest_next_topic()
        
        if not topic:
            print("🚀 EvolvAI has completed its study.")
            break
        
        brain.learn_topic(topic)

        print("\n📌 EvolvAI has reached a learning checkpoint.")
        choice = input("Continue auto-learning (yes), select a field (no), or stop (exit)? ").strip().lower()
        
        if choice == "no":
            topic = input("Enter a topic to guide EvolvAI (or press Enter to stop): ").strip()
            if not topic:
                print("🚀 EvolvAI has completed its study.")
                break
            brain.learn_topic(topic)

        elif choice == "exit":
            print("🚀 EvolvAI has completed its study.")
            break
        
        time.sleep(1)

if __name__ == "__main__":
    main()
