import { useEffect, useState } from "react";
import Link from "next/link";
import { FiEdit, FiTrash2, FiXCircle } from "react-icons/fi"; // Import icons
import { useRouter } from "next/router";

const TabButton = ({ label, isActive, onClick }) => (
  <button
    className={`inline-block py-2.5 px-5 rounded-md text-sm font-medium transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 ${
      isActive
        ? "bg-indigo-500 text-white shadow-md"
        : "bg-white text-gray-700 hover:bg-gray-100"
    }`}
    onClick={onClick}
  >
    {label}
  </button>
);

const ActionButton = ({ icon, onClick, colorClass }) => (
  <button
    onClick={onClick}
    className={`inline-flex items-center justify-center rounded-md p-2 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-${colorClass}-500 hover:bg-${colorClass}-50`}
  >
    <span className={`text-lg text-${colorClass}-500`}>{icon}</span>
  </button>
);

export default function AllPosts() {
  const [activeTab, setActiveTab] = useState("Publish"); // Default to Publish
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const router = useRouter();

  useEffect(() => {
    fetchArticles();
  }, [activeTab]);

  const fetchArticles = async () => {
    setLoading(true);
    setError(null);
    try {
      let url = "http://localhost:8001/api/articles";
      if (activeTab !== "All") {
        url = `http://localhost:8001/api/articles/search?q=${activeTab.toLowerCase()}`;
      }

      const res = await fetch(url, {
        headers: {
          "Content-Type": "application/json",
          // Tambahkan otentikasi jika diperlukan
        },
      });

      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      const data = await res.json();
      setArticles(data);
    } catch (err) {
      setError(err.message);
      setArticles([]);
    } finally {
      setLoading(false);
    }
  };

  const handleTrash = async (id) => {
    if (activeTab === "Trash") {
      // Jika di tab Trash, lakukan penghapusan permanen
      if (
        window.confirm(
          "PERHATIAN: Artikel ini akan dihapus PERMANENLY. Anda yakin ingin melanjutkan?"
        )
      ) {
        try {
          const res = await fetch(`http://localhost:8001/api/articles/${id}`, {
            // Menggunakan endpoint DELETE
            method: "DELETE",
            headers: {
              "Content-Type": "application/json",
              // Tambahkan otentikasi jika diperlukan
            },
          });

          if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
          }

          fetchArticles(); // Re-fetch articles setelah dihapus
          alert("Artikel berhasil dihapus secara permanen.");
        } catch (err) {
          setError(err.message);
          alert(`Gagal menghapus artikel: ${err.message}`);
        }
      }
    } else {
      // Jika tidak di tab Trash, lakukan soft delete (pindahkan ke Trash)
      if (
        window.confirm(
          "Apakah Anda yakin ingin memindahkan artikel ini ke Trash?"
        )
      ) {
        try {
          const res = await fetch(
            `http://localhost:8001/api/articles/${id}/trash`,
            {
              // Menggunakan endpoint soft delete
              method: "PUT",
              headers: {
                "Content-Type": "application/json",
                // Tambahkan otentikasi jika diperlukan
              },
            }
          );

          if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
          }

          fetchArticles(); // Re-fetch articles setelah di-trash
          alert("Artikel berhasil dipindahkan ke Trash.");
        } catch (err) {
          setError(err.message);
          alert(`Gagal memindahkan ke Trash: ${err.message}`);
        }
      }
    }
  };

  return (
    <div className="w-full">
      <div className="flex flex-col gap-3 !m-10 !p-8 bg-white shadow-lg rounded-xl">
        <div className="flex justify-between">
          <h2 className="text-3xl  font-semibold text-gray-800 mb-8">
            Postingan Article
          </h2>
          <Link
            href="/articles/add-article"
            className="!px-8 !bg-indigo-100 !text-indigo-500 !rounded-full"
          >
            Add Article
          </Link>
        </div>

        <div className="flex gap-3 mb-8 space-x-3">
          <TabButton
            label="Published"
            isActive={activeTab === "Publish"}
            onClick={() => setActiveTab("Publish")}
          />
          <TabButton
            label="Drafts"
            isActive={activeTab === "draft"}
            onClick={() => setActiveTab("draft")}
          />
          <TabButton
            label="Trashed"
            isActive={activeTab === "Trash"}
            onClick={() => setActiveTab("Trash")}
          />
          <TabButton
            label="All"
            isActive={activeTab === "All"}
            onClick={() => setActiveTab("All")}
          />
        </div>

        {loading && (
          <div className="text-center py-6 text-gray-500">
            Memuat postingan...
          </div>
        )}
        {error && (
          <div className="text-center py-6 text-red-500">Error: {error}</div>
        )}

        {!loading && !error && (
          <div className="overflow-x-auto">
            <table className="min-w-full leading-normal shadow-md rounded-lg">
              <thead className=" bg-gray-50 border-b border-gray-200">
                <tr>
                  <th className="!px-6 !py-3 text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">
                    Judul
                  </th>
                  <th className="!px-6 !py-3  text-left text-sm font-semibold text-gray-700 uppercase tracking-wider">
                    Kategori
                  </th>
                  <th className="!px-6 !py-3  text-right text-sm font-semibold text-gray-700 uppercase tracking-wider">
                    Aksi
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white">
                {articles && articles.length > 0 ? (
                  articles.map((article) => (
                    <tr
                      key={article.id}
                      className="hover:bg-gray-50 transition-colors duration-200"
                    >
                      <td className="!px-6 !py-4  whitespace-nowrap text-sm font-medium text-gray-900">
                        <Link
                          href={`/articles/${article.id}`}
                          key={`${article.id}`}
                        >
                          {article.title}
                        </Link>
                      </td>
                      <td className="!px-6 !py-4 whitespace-nowrap text-sm text-gray-700">
                        <span className="inline-block bg-indigo-100 text-indigo-500 rounded-full !px-3 !py-1 font-semibold text-sm">
                          {article.category}
                        </span>
                      </td>
                      <td className="!px-6 !py-4 whitespace-nowrap text-right text-sm font-medium">
                        <div className="space-x-3">
                          <Link href={`/articles/edit-article/${article.id}`}>
                            <ActionButton icon={<FiEdit />} colorClass="blue" />
                          </Link>
                          {activeTab === "Trash" ? (
                            <ActionButton
                              icon={<FiXCircle />} // Ganti ikon menjadi ikon delete permanen
                              onClick={() => handleTrash(article.id)}
                              colorClass="red"
                              key={`delete-${article.id}`}
                            />
                          ) : (
                            <ActionButton
                              icon={<FiTrash2 />}
                              onClick={() => handleTrash(article.id)}
                              colorClass="red"
                              key={`trash-${article.id}`}
                            />
                          )}
                        </div>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr key="no-articles">
                    <td
                      className="px-6 py-4 whitespace-nowrap text-center text-sm text-gray-500"
                      colSpan="3"
                    >
                      Tidak ada postingan dengan status {activeTab}
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}
