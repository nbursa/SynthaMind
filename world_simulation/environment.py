import random
from core_intelligence.sensory_specialization import SensorySpecialization
from core_intelligence.knowledge_integration import KnowledgeIntegration

class AIWorld:
    """
    A simple 2D grid world where EvolvAI moves, remembers visited locations, 
    and stops learning once everything is explored.
    """

    KNOWLEDGE_CATEGORIES = ["medicine", "mechanics", "physics", "biology", "mathematics"]

    def __init__(self, width=10, height=10):
        self.width = width
        self.height = height
        self.grid = self._generate_world()  # ✅ FIXED: Ensure this method exists
        self.visited_positions = set()  # Tracks visited locations
        self.ai_position = self._place_ai()
        self.specialization_engine = SensorySpecialization()  # AI specialization engine
        self.knowledge_integration = KnowledgeIntegration(self.specialization_engine)  # NEW: Connect knowledge system

    def _generate_world(self):
        """ ✅ FIXED: Generates a 2D grid with random objects categorized by knowledge fields. """
        return [
            [random.choice([None, None, (random.choice(self.KNOWLEDGE_CATEGORIES), "O")]) for _ in range(self.width)]
            for _ in range(self.height)
        ]

    def _place_ai(self):
        """ Places the AI at a random, empty starting position """
        while True:
            x, y = random.randint(0, self.width - 1), random.randint(0, self.height - 1)
            if self.grid[y][x] is None:  # Ensure AI doesn't start on an object
                self.visited_positions.add((x, y))  # Mark initial position as visited
                return x, y

    def get_valid_moves(self):
        """ Returns possible movements AI can take, prioritizing unexplored areas. """
        x, y = self.ai_position
        moves = []
        potential_moves = {
            "LEFT": (x - 1, y),
            "RIGHT": (x + 1, y),
            "UP": (x, y - 1),
            "DOWN": (x, y + 1)
        }

        # Prioritize unvisited locations
        unexplored_moves = [move for move, pos in potential_moves.items()
                            if pos not in self.visited_positions and 0 <= pos[0] < self.width and 0 <= pos[1] < self.height]
        
        return unexplored_moves if unexplored_moves else list(potential_moves.keys())  # Prioritize unexplored, else move randomly

    def execute_move(self, move):
        """ Moves AI and checks for stimuli, integrating related knowledge. """
        x, y = self.ai_position

        if move == "LEFT" and x > 0:
            x -= 1
        elif move == "RIGHT" and x < self.width - 1:
            x += 1
        elif move == "UP" and y > 0:
            y -= 1
        elif move == "DOWN" and y < self.height - 1:
            y += 1
        else:
            return None  # Invalid move, do nothing

        self.ai_position = (x, y)
        self.visited_positions.add((x, y))  # Track visited locations

        # Check if AI found an object
        stimulus = self.grid[y][x]
        if stimulus:
            category, _ = stimulus  # Extract category of object
            self.grid[y][x] = None  # Object is "consumed"
            self.specialization_engine.update_specialization(category)  # AI learns from discovery
            self.knowledge_integration.integrate_knowledge(category)  # NEW: Link related fields
            print(f"EvolvAI found something related to {category}!")
            self.knowledge_integration.print_knowledge()
            return category

        return None  # No new information

    def has_fully_explored(self):
        """ Checks if EvolvAI has visited all possible locations in the world. """
        total_possible_positions = self.width * self.height  # All grid cells
        return len(self.visited_positions) >= total_possible_positions
