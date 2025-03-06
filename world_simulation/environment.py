import random
from core_intelligence.sensory_specialization import SensorySpecialization

class AIWorld:
    """
    A simple 2D grid world where EvolvAI can move and discover objects linked to different knowledge categories.
    """

    KNOWLEDGE_CATEGORIES = ["medicine", "mechanics", "physics", "biology", "mathematics"]

    def __init__(self, width=10, height=10):
        self.width = width
        self.height = height
        self.grid = self._generate_world()
        self.ai_position = self._place_ai()
        self.specialization_engine = SensorySpecialization()  # Connects to specialization engine

    def _generate_world(self):
        """ Generates a 2D grid with random objects categorized by knowledge fields. """
        return [
            [random.choice([None, None, (random.choice(self.KNOWLEDGE_CATEGORIES), "O")]) for _ in range(self.width)]
            for _ in range(self.height)
        ]

    def _place_ai(self):
        """ Places the AI at a random, empty starting position """
        while True:
            x, y = random.randint(0, self.width - 1), random.randint(0, self.height - 1)
            if self.grid[y][x] is None:  # Ensure AI doesn't start on an object
                return x, y

    def get_valid_moves(self):
        """ Returns possible movements AI can take (avoiding walls) """
        x, y = self.ai_position
        moves = []
        if x > 0: moves.append("LEFT")
        if x < self.width - 1: moves.append("RIGHT")
        if y > 0: moves.append("UP")
        if y < self.height - 1: moves.append("DOWN")
        return moves

    def execute_move(self, move):
        """ Moves AI in the selected direction and checks for stimuli """
        x, y = self.ai_position

        if move == "LEFT": x -= 1
        elif move == "RIGHT": x += 1
        elif move == "UP": y -= 1
        elif move == "DOWN": y += 1

        self.ai_position = (x, y)

        # Check if AI found an object
        stimulus = self.grid[y][x]
        if stimulus:
            category, _ = stimulus  # Extract category of object
            self.grid[y][x] = None  # Object is "consumed"
            self.specialization_engine.update_specialization(category)  # AI learns from discovery
            print(f"EvolvAI found something related to {category}!")
            self.specialization_engine.print_specialization()
            return category

        return None  # No new information
