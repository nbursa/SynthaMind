class EvolvAIInstincts:
    """
    Defines the immutable core instincts for EvolvAI.
    These instincts never change and define its fundamental behavior.
    """

    _instance = None  # Singleton instance

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(EvolvAIInstincts, cls).__new__(cls)
            cls._instance._initialize_instincts()
        return cls._instance

    def _initialize_instincts(self):
        """ Initializes fixed instinct values """
        
        # Curiosity: EvolvAI will seek new knowledge if entropy is low
        self.curiosity_threshold = 0.3  # Below this, AI gets "bored" and seeks novelty

        # Logic & Pattern Recognition: AI must detect order in data
        self.pattern_detection_threshold = 0.7  # Above this, AI recognizes structure

        # Sensory Limits: AI can only process a fixed amount of data per cycle
        self.sensory_limit = 50  # Number of input elements AI can process per iteration

        # Self-Preservation: AI prevents actions that could cause logical contradictions
        self.self_preservation_enabled = True  # AI avoids infinite loops or data collapse

        # Memory Instinct: AI stores meaningful patterns but forgets irrelevant noise
        self.memory_decay_rate = 0.01  # Gradual forgetting of older, less relevant patterns

    def get_instincts(self):
        """ Returns the immutable instinct values """
        return {
            "curiosity_threshold": self.curiosity_threshold,
            "pattern_detection_threshold": self.pattern_detection_threshold,
            "sensory_limit": self.sensory_limit,
            "self_preservation_enabled": self.self_preservation_enabled,
            "memory_decay_rate": self.memory_decay_rate
        }
