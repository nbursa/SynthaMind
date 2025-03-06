class KnowledgeIntegration:
    """
    EvolvAI starts linking knowledge fields by recognizing related disciplines.
    """

    RELATED_FIELDS = {
        "medicine": ["biology"],
        "biology": ["medicine", "physics"],
        "physics": ["mechanics", "mathematics"],
        "mechanics": ["physics"],
        "mathematics": ["physics"]
    }

    def __init__(self, specialization_engine):
        self.specialization_engine = specialization_engine  # Connect to existing specialization system

    def integrate_knowledge(self, category):
        """ When EvolvAI learns a field, related fields also gain a small boost. """
        if category in self.RELATED_FIELDS:
            for related_field in self.RELATED_FIELDS[category]:
                self.specialization_engine.update_specialization(related_field, boost=0.05)  # Small boost for related fields

    def print_knowledge(self):
        """ Debugging: Show EvolvAI’s specialization after integration. """
        self.specialization_engine.print_specialization()
