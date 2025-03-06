class SensorySpecialization:
    """
    Tracks what EvolvAI learns over time and adjusts specialization weights.
    EvolvAI becomes more interested in fields it encounters often.
    """

    def __init__(self):
        # Initialize interest levels for different knowledge areas
        self.specialization_weights = {
            "medicine": 1.0,
            "mechanics": 1.0,
            "physics": 1.0,
            "biology": 1.0,
            "mathematics": 1.0
        }

    def update_specialization(self, category):
        """ Increase interest in a specific category when encountered. """
        if category in self.specialization_weights:
            self.specialization_weights[category] += 0.1  # Small weight increase
            self.normalize_weights()

    def normalize_weights(self):
        """ Ensures AI doesn’t over-specialize too quickly. """
        total = sum(self.specialization_weights.values())
        for key in self.specialization_weights:
            self.specialization_weights[key] /= total  # Keep values balanced

    def get_specialization_focus(self):
        """ Returns the field EvolvAI is currently most interested in. """
        return max(self.specialization_weights, key=self.specialization_weights.get)

    def print_specialization(self):
        """ Debugging: Show EvolvAI’s current specialization levels. """
        print("Current specialization weights:", self.specialization_weights)
