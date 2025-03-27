# **Contributing to SynthaMind**

Thank you for considering contributing to **SynthaMind**! 🚀

## **Contribution Guidelines**

- **Fork the repository** and submit a **Pull Request** with a detailed description.
- Ensure your contributions align with the **AGPLv3 license** (all changes remain open-source).
- Clearly document your code and follow the project structure.

## **Rules for Contributions**

✔ All code must be properly structured and tested.  
✔ Any changes must be in line with the **AI framework's core vision**.  
✔ Contributors must **credit the original creator, Nenad Bursać**, when referring to the project's ideas.

## **How to Get Started?**

1. **Clone the repository**
   ```bash
   git clone https://github.com/nbursa/SynthaMind.git
   ```
2. **Create a new branch for your changes**
   ```bash
   git checkout -b feature-branch-name
   ```
3. **Make your changes and commit them**
   ```bash
   git add .
   git commit -m "Describe your changes"
   ```
4. **Push your changes to GitHub**
   ```bash
   git push origin feature-branch-name
   ```
5. **Submit a pull request for review**

## **Running ChromaDB Locally**

To run **ChromaDB** locally for storing and retrieving vector data, follow these steps:

1. **Ensure Python 3.7+ and pip are installed** on your machine.
2. **Install ChromaDB** (if not already installed) by running:
   ```bash
   pip install chromadb
   ```
3. **Start ChromaDB server**:

   ```bash
   uvicorn chromadb.app:app --host 127.0.0.1 --port 8000
   ```

   This will start the ChromaDB server on `http://127.0.0.1:8000`. Ensure the server is running before interacting with it.

4. **Verify server status**:
   You can verify if ChromaDB is running by navigating to `http://127.0.0.1:8000` in your browser or using a tool like `curl`:
   ```bash
   curl http://127.0.0.1:8000/api/v1
   ```

## **Contact**

For major feature proposals or discussions about collaborations, visit:  
**Website:** [https://nenadbursac.com/contact](https://nenadbursac.com/contact)
