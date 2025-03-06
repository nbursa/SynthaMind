import random
from core_intelligence.instinct_processor import EvolvAIInstinctProcessor
from world_simulation.environment import AIWorld
import time

# Initialize the world and AI
world = AIWorld(width=10, height=10)
ai = EvolvAIInstinctProcessor(world)

print("EvolvAI is starting...")

while True:
    if world.has_fully_explored():
        print("EvolvAI has fully explored the world. Exploration complete.")
        break

    valid_moves = world.get_valid_moves()
    if not valid_moves:
        print("EvolvAI has no valid moves left. Exploration complete.")
        break

    random_move = random.choice(valid_moves)
    new_stimulus = world.execute_move(random_move)

    if new_stimulus:
        print(f"EvolvAI found something related to {new_stimulus}!")
    else:
        print("EvolvAI is bored! Exploring the world...")

    time.sleep(1)  # Simulate real-time exploration
