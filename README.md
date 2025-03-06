### **🧠 EvolvAI**

EvolvAI is an **adaptive knowledge exploration system** that simulates **self-learning behavior** by dynamically retrieving and storing knowledge. It can explore topics based on **specialization**, **user guidance**, or **self-directed curiosity**.

While **not a true AI agent (yet)**, EvolvAI serves as a **foundation for an AI-driven learning framework**, where future versions will integrate **machine learning, reinforcement learning, and intelligent decision-making**.

---

## **🎯 Goal**

The ultimate goal of EvolvAI is to create an **autonomous AI system** that can:

- **Explore knowledge dynamically** across different domains.
- **Self-specialize** based on discovered knowledge and user input.
- **Store, refine, and connect knowledge** to build an intelligent knowledge base.
- **Make learning decisions intelligently** instead of following fixed rules.

Right now, EvolvAI operates as a **structured, rule-based knowledge explorer** that retrieves and integrates knowledge **without AI-driven inference**.

---

## **⚙️ Current Features**

✅ **Self-Guided Knowledge Exploration**

- EvolvAI explores knowledge based on user-selected fields (e.g., AI, medicine, physics).
- Auto-explores related topics or follows user-defined paths.

✅ **Memory Retention**

- Prevents repeated learning of the same topics.
- Stores learned knowledge for future use.

✅ **Learning Checkpoints**

- Periodically pauses learning to ask if it should **continue automatically**, **switch focus**, or **stop**.

✅ **User & Auto-Guided Specialization**

- EvolvAI can **auto-learn** or **be guided** to focus on specific topics.

✅ **Wikipedia Knowledge Integration**

- Fetches structured summaries from Wikipedia for knowledge acquisition.

---

## **🚀 Next Steps (Future AI Integration)**

EvolvAI is **not yet an AI agent**, but the roadmap includes:

1. **AI-Based Topic Selection**

   - Use **semantic search and embeddings** to **dynamically retrieve related topics** instead of basic keyword matching.

2. **AI-Driven Memory & Storage**

   - Replace the JSON knowledge base with **vector embeddings** (e.g., FAISS, ChromaDB) for **intelligent retrieval**.

3. **Reinforcement Learning for Decision Making**

   - Enable EvolvAI to **evaluate what to learn next** based on **novelty, depth, and relevance**.

4. **Natural Language Processing (NLP) Integration**

   - Use **LLMs (GPT, BERT, LLaMA)** to **generate insights, summarize data**, and **explain knowledge**.

5. **Self-Optimization**
   - EvolvAI should **adjust its own learning strategies over time**, improving **exploration efficiency**.

---

## **🛠️ How to Use**

1️⃣ **Run EvolvAI**

```bash
python main.py
```

2️⃣ **Choose a specialization field**

- Example: `AI`, `Medicine`, `Physics`  
  3️⃣ **Select Auto or Guided Exploration**
- Auto: EvolvAI explores on its own.
- Guided: User inputs topics to guide learning.  
  4️⃣ **Periodically, EvolvAI pauses at checkpoints**
- Options: **Continue auto-learning**, **switch focus**, or **stop**.
