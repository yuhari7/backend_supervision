import { useEffect, useState } from "react";
import { useRouter } from "next/router";

export default function EditArticle() {
  const router = useRouter();
  const { id } = router.query; // Mengambil id artikel dari URL
  const [article, setArticle] = useState({
    title: "",
    content: "",
    category: "",
    status: "Draft",
  });
  const [errors, setErrors] = useState({});
  const [loading, setLoading] = useState(false);
  const [apiError, setApiError] = useState(null);

  // Fetch artikel berdasarkan id saat halaman dimuat
  useEffect(() => {
    if (id) {
      fetchArticle();
    }
  }, [id]);

  const fetchArticle = async () => {
    setLoading(true);
    setApiError(null);
    try {
      const res = await fetch(`http://localhost:8001/api/articles/${id}`);
      if (!res.ok) {
        const errorData = await res.json();
        setApiError(errorData?.error || `Error: Status ${res.status}`);
        return;
      }

      const data = await res.json();
      setArticle(data);
    } catch (err) {
      setApiError(`Error: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setArticle({ ...article, [name]: value });
    if (errors[name]) {
      setErrors({ ...errors, [name]: "" });
    }
  };

  // Validasi input form
  const validate = () => {
    let isValid = true;
    const newErrors = {};

    if (!article.title.trim() || article.title.length < 20) {
      newErrors.title = "Judul minimal 20 karakter.";
      isValid = false;
    }

    if (!article.content.trim() || article.content.length < 200) {
      newErrors.content = "Konten minimal 200 karakter.";
      isValid = false;
    }

    if (!article.category.trim() || article.category.length < 3) {
      newErrors.category = "Kategori minimal 3 karakter.";
      isValid = false;
    }

    setErrors(newErrors);
    return isValid;
  };

  // Handle submit form untuk update artikel
  const handleUpdate = async () => {
    if (!validate()) {
      return;
    }

    setLoading(true);
    setApiError(null);
    try {
      const res = await fetch(`http://localhost:8001/api/articles/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(article),
      });

      if (!res.ok) {
        const errorData = await res.json();
        setApiError(errorData?.error || `Error: Status ${res.status}`);
        return;
      }

      alert("Artikel berhasil diperbarui.");
      router.push("/articles");
    } catch (err) {
      setApiError(`Error: ${err.message}`);
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

  if (apiError)
    return (
      <div className="container mx-auto p-8 bg-white shadow-lg rounded-xl">
        <p className="text-center text-red-500">
          Terjadi kesalahan: {apiError}
        </p>
      </div>
    );

  return (
    <div className="w-full">
      <div className="flex flex-col gap-3 !m-10 !p-8 bg-white shadow-lg rounded-xl">
        <h2 className="text-3xl font-semibold text-gray-800 mb-8">
          Edit Artikel
        </h2>
        <form className="flex flex-col gap-3  space-y-8">
          <div>
            <label
              htmlFor="title"
              className="block text-sm font-medium text-gray-700"
            >
              Judul:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={article.title}
              onChange={handleChange}
              className={`shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md ${
                errors.title ? "border-red-500" : ""
              }`}
            />
            {errors.title && (
              <p className="text-red-500 text-xs italic">{errors.title}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="content"
              className="block text-sm font-medium text-gray-700"
            >
              Konten:
            </label>
            <textarea
              id="content"
              name="content"
              value={article.content}
              onChange={handleChange}
              rows="8"
              className={`shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md ${
                errors.content ? "border-red-500" : ""
              }`}
            />
            {errors.content && (
              <p className="text-red-500 text-xs italic">{errors.content}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="category"
              className="block text-sm font-medium text-gray-700"
            >
              Kategori:
            </label>
            <input
              type="text"
              id="category"
              name="category"
              value={article.category}
              onChange={handleChange}
              className={`shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md ${
                errors.category ? "border-red-500" : ""
              }`}
            />
            {errors.category && (
              <p className="text-red-500 text-xs italic">{errors.category}</p>
            )}
          </div>
          <div className="flex gap-4 space-x-4">
            <button
              type="button"
              className="!bg-indigo-600 !hover:bg-indigo-700 !text-white !font-bold !py-3 !px-6 !rounded-md focus:outline-none focus:shadow-outline"
              onClick={handleUpdate}
            >
              Simpan Perubahan
            </button>
            <button
              type="button"
              className="!bg-gray-300 !hover:bg-gray-400 !text-gray-700 !font-bold !py-3 !px-6 !rounded-md focus:outline-none focus:shadow-outline"
              onClick={() => router.push("/articles")}
            >
              Batal
            </button>
          </div>
          {loading && (
            <p className="text-center text-gray-500">Menyimpan perubahan...</p>
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
