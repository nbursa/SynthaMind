from core_intelligence.knowledge_integration import KnowledgeIntegration
from core_intelligence.sensory_specialization import SensorySpecialization

class AIWorld:
    def __init__(self, width=10, height=10):
        self.width = width
        self.height = height
        self.grid = self._generate_world()
        self.visited_positions = set()
        self.ai_position = self._place_ai()
        self.specialization_engine = SensorySpecialization()
        self.knowledge_integration = KnowledgeIntegration(self.specialization_engine)  # NEW: Connect knowledge system

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
            return None

        self.ai_position = (x, y)
        self.visited_positions.add((x, y))

        stimulus = self.grid[y][x]
        if stimulus:
            category, _ = stimulus
            self.grid[y][x] = None
            self.specialization_engine.update_specialization(category)
            self.knowledge_integration.integrate_knowledge(category)  # NEW: Link related fields
            print(f"EvolvAI found something related to {category}!")
            self.knowledge_integration.print_knowledge()
            return category

        return None
