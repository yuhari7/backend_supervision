import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import Link from "next/link";

export default function ArticleDetail() {
  const router = useRouter();
  const { id } = router.query;
  const [article, setArticle] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (id) {
      fetchArticle();
    }
  }, [id]);

  const fetchArticle = async () => {
    setLoading(true);
    setError(null);
    setArticle(null); // Reset artikel saat ID berubah
    try {
      const res = await fetch(`http://localhost:8001/api/articles/${id}`, {
        headers: {
          "Content-Type": "application/json",
          // Tambahkan otentikasi jika diperlukan
        },
      });

      if (!res.ok) {
        const errorData = await res.json();
        if (errorData && errorData.message) {
          setError(errorData.message);
        } else {
          setError(`Gagal memuat artikel: Status ${res.status}`);
        }
        return;
      }

      const data = await res.json();
      setArticle(data);
    } catch (err) {
      setError(`Gagal menghubungi server: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  if (loading)
    return (
      <div className="container mx-auto p-8 bg-white shadow-lg rounded-xl">
        <p className="text-center text-gray-500">Memuat artikel...</p>
      </div>
    );
  if (error)
    return (
      <div className="container mx-auto p-8 bg-white shadow-lg rounded-xl">
        <p className="text-center text-red-500">Error: {error}</p>
      </div>
    );
  if (!article)
    return (
      <div className="container mx-auto p-8 bg-white shadow-lg rounded-xl">
        <p className="text-center text-gray-500">Artikel tidak ditemukan.</p>
      </div>
    );

  return (
    <div className="w-full">
      <div className="flex flex-col gap-3 !m-10 !p-8 bg-white shadow-lg rounded-xl">
        <h1 className="text-4xl font-semibold text-indigo-700 mb-6">
          {article.title}
        </h1>
        <div className="mb-4 text-gray-600">
          Kategori: <span className="font-semibold">{article.category}</span>
        </div>
        <div className="mb-6 text-gray-800 whitespace-pre-line">
          {article.content}
        </div>
        <div className="mt-8">
          <Link href="/articles" className="!text-indigo-500 !hover:underline">
            Kembali ke Daftar Artikel
          </Link>
        </div>
      </div>
    </div>
  );
}
