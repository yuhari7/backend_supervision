import { useState } from "react";
import { useRouter } from "next/router";

export default function AddNew() {
  const router = useRouter();
  const [article, setArticle] = useState({
    Title: "",
    Content: "",
    Category: "",
    Status: "draft",
  });
  const [errors, setErrors] = useState({});
  const [loading, setLoading] = useState(false);
  const [apiError, setApiError] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setArticle({ ...article, [name]: value });
    // Hapus error terkait saat nilai input berubah
    if (errors[name]) {
      setErrors({ ...errors, [name]: "" });
    }
  };

  const validate = () => {
    let isValid = true;
    const newErrors = {};

    if (!article.Title.trim()) {
      newErrors.Title = "Judul harus diisi.";
      isValid = false;
    } else if (article.Title.length > 255) {
      newErrors.Title = "Judul tidak boleh lebih dari 255 karakter.";
      isValid = false;
    }

    if (!article.Content.trim()) {
      newErrors.Content = "Konten harus diisi.";
      isValid = false;
    }

    if (!article.Category.trim()) {
      newErrors.Category = "Kategori harus diisi.";
      isValid = false;
    } else if (article.Category.length > 100) {
      newErrors.Category = "Kategori tidak boleh lebih dari 100 karakter.";
      isValid = false;
    }

    setErrors(newErrors);
    return isValid;
  };

  const handleSave = async (status) => {
    if (!validate()) {
      return;
    }

    setLoading(true);
    setApiError(null);
    try {
      const res = await fetch("http://localhost:8001/api/articles", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          // Tambahkan otentikasi jika diperlukan
        },
        body: JSON.stringify({ ...article, Status: status }),
      });

      if (!res.ok) {
        const errorData = await res.json();
        if (errorData && errorData.message) {
          setApiError(errorData.message);
        } else {
          setApiError(
            `Terjadi kesalahan saat menyimpan artikel: Status ${res.status}`
          );
        }
        return;
      }

      alert(
        `Artikel berhasil di${
          status === "Publish" ? "publikasikan" : "simpan sebagai draft"
        }.`
      );
      router.push("/articles");
    } catch (err) {
      setApiError(`Gagal menghubungi server: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full">
      <div className="flex flex-col gap-3 !m-10 !p-8 bg-white shadow-lg rounded-xl">
        <h2 className="text-3xl font-semibold text-gray-800 mb-8">
          Buat Artikel Baru
        </h2>
        <form className="flex flex-col gap-3 space-y-8">
          <div>
            <label
              htmlFor="Title"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Judul:
            </label>
            <input
              type="text"
              id="Title"
              name="Title"
              value={article.Title}
              onChange={handleChange}
              className={`shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 ${
                errors.Title ? "border-red-500" : ""
              }`}
            />
            {errors.Title && (
              <p className="text-red-500 text-xs italic">{errors.Title}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="Content"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Konten:
            </label>
            <textarea
              id="Content"
              name="Content"
              value={article.Content}
              onChange={handleChange}
              rows="10"
              className={`shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 ${
                errors.Content ? "border-red-500" : ""
              }`}
            />
            {errors.Content && (
              <p className="text-red-500 text-xs italic">{errors.Content}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="Category"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Kategori:
            </label>
            <input
              type="text"
              id="Category"
              name="Category"
              value={article.Category}
              onChange={handleChange}
              className={`shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 ${
                errors.Category ? "border-red-500" : ""
              }`}
            />
            {errors.Category && (
              <p className="text-red-500 text-xs italic">{errors.Category}</p>
            )}
          </div>
          <div className="flex gap-4 space-x-4">
            <button
              type="button"
              className="!bg-indigo-600 !hover:bg-indigo-700 !text-white !font-bold !py-3 !px-6 !rounded-md focus:outline-none focus:shadow-outline"
              onClick={() => handleSave("Publish")}
            >
              Publish
            </button>
            <button
              type="button"
              className="!bg-gray-300 !hover:bg-gray-400 !text-gray-700 !font-bold !py-3 !px-6 !rounded-md focus:outline-none focus:shadow-outline"
              onClick={() => handleSave("draft")}
            >
              Simpan Draft
            </button>
          </div>
          {loading && (
            <p className="text-center text-gray-500">Menyimpan artikel...</p>
          )}
          {apiError && (
            <p className="text-center text-red-500">
              Terjadi kesalahan: {apiError}
            </p>
          )}
        </form>
      </div>
    </div>
  );
}
