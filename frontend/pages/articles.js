import { useEffect, useState } from "react";

export default function Articles() {
  const [articles, setArticles] = useState([]);

  useEffect(() => {
    const fetchArticles = async () => {
      const token = localStorage.getItem("access_token");

      const response = await fetch("http://localhost:8001/api/articles", {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = await response.json();

      if (response.ok) {
        setArticles(data);
      } else {
        alert(data.error || "Failed to fetch articles");
      }
    };

    fetchArticles();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100">
      <div className="container mx-auto p-4">
        <h2 className="text-2xl font-bold mb-6">Articles</h2>
        <div className="space-y-4">
          {articles.map((article) => (
            <div key={article.id} className="p-4 bg-white shadow-md rounded">
              <h3 className="text-xl font-semibold">{article.title}</h3>
              <p>{article.content}</p>
              <span className="text-gray-500">{article.category}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
