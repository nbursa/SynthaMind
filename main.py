from world_simulation.environment import AIWorld
from core_intelligence.instinct_processor import EvolvAIInstinctProcessor
import time

# Initialize the world and AI
world = AIWorld(width=10, height=10)  # Create a simple 10x10 grid world
ai = EvolvAIInstinctProcessor(world)

print("EvolvAI is starting...")

while True:
    if world.has_fully_explored():
        print("🚀 EvolvAI has fully explored its world! Learning is complete.")
        break  # Stop execution

    if ai.check_boredom():
        print("EvolvAI is bored! Exploring the world...")
        ai.explore_world()

    time.sleep(1)  # Simulate real-time updates
