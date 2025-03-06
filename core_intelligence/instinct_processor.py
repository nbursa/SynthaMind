from core_intelligence.instincts import EvolvAIInstincts
from world_simulation.environment import AIWorld
import numpy as np
import time
import random

class EvolvAIInstinctProcessor:
    """
    Processes AI instincts and applies them to real-time learning and perception.
    """

    def __init__(self, world):
        self.instincts = EvolvAIInstincts()
        self.world = world  # AI interacts with this world
        self.last_stimulation_time = time.time()  # Tracks last input time

    def enforce_sensory_limits(self, data):
        """ Limits AI perception to prevent overwhelming input """
        sensory_limit = self.instincts.sensory_limit
        return data[:sensory_limit] if len(data) > sensory_limit else data

    def check_boredom(self):
        """ EvolvAI gets bored if no meaningful input occurs for 10 seconds """
        return time.time() - self.last_stimulation_time >= 10

    def update_stimulation_time(self):
        """ Resets boredom timer when new input is received """
        self.last_stimulation_time = time.time()

    def explore_world(self):
        """ AI moves randomly in its environment to seek new stimuli when bored.
            If everything is explored, it stops learning.
        """
        if self.world.has_fully_explored():
            print("🚀 EvolvAI has fully explored its world! Stopping learning.")
            return None  # Stop further exploration

        possible_moves = self.world.get_valid_moves()
        if not possible_moves:
            return None  # No available moves

        random_move = random.choice(possible_moves)  # Pick a random action
        new_stimulus = self.world.execute_move(random_move)  # AI moves & interacts

        if new_stimulus:
            self.update_stimulation_time()  # Reset boredom if new information is found
            return new_stimulus

        return None  # Continue exploring if nothing new is found
