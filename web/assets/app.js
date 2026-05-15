const state = {
  token: localStorage.getItem("product_manager_token") || "",
  products: [],
  users: [],
  mode: "login",
  editingId: null,
  view: "products",
  productPage: 1,
  productLimit: 10,
  productTotal: 0,
  productTotalPages: 1,
};

const els = {
  authPanel: document.querySelector("#authPanel"),
  dashboard: document.querySelector("#dashboard"),
  sessionState: document.querySelector("#sessionState"),
  logoutButton: document.querySelector("#logoutButton"),
  authForm: document.querySelector("#authForm"),
  loginTab: document.querySelector("#loginTab"),
  registerTab: document.querySelector("#registerTab"),
  authSubmit: document.querySelector("#authSubmit"),
  authMessage: document.querySelector("#authMessage"),
  email: document.querySelector("#email"),
  password: document.querySelector("#password"),
  productsViewTab: document.querySelector("#productsViewTab"),
  usersViewTab: document.querySelector("#usersViewTab"),
  viewTitle: document.querySelector("#viewTitle"),
  navProductCount: document.querySelector("#navProductCount"),
  navUserCount: document.querySelector("#navUserCount"),
  productsView: document.querySelector("#productsView"),
  usersView: document.querySelector("#usersView"),
  openProductDrawer: document.querySelector("#openProductDrawer"),
  closeProductDrawer: document.querySelector("#closeProductDrawer"),
  productDrawer: document.querySelector("#productDrawer"),
  drawerBackdrop: document.querySelector("#drawerBackdrop"),
  productForm: document.querySelector("#productForm"),
  formTitle: document.querySelector("#formTitle"),
  cancelEdit: document.querySelector("#cancelEdit"),
  productName: document.querySelector("#productName"),
  productDescription: document.querySelector("#productDescription"),
  productPrice: document.querySelector("#productPrice"),
  productStock: document.querySelector("#productStock"),
  saveProduct: document.querySelector("#saveProduct"),
  addExamples: document.querySelector("#addExamples"),
  productMessage: document.querySelector("#productMessage"),
  productRows: document.querySelector("#productRows"),
  emptyState: document.querySelector("#emptyState"),
  productPageInfo: document.querySelector("#productPageInfo"),
  prevProductsPage: document.querySelector("#prevProductsPage"),
  nextProductsPage: document.querySelector("#nextProductsPage"),
  userRows: document.querySelector("#userRows"),
  emptyUsersState: document.querySelector("#emptyUsersState"),
  searchUsers: document.querySelector("#searchUsers"),
  searchProducts: document.querySelector("#searchProducts"),
  stockFilter: document.querySelector("#stockFilter"),
  totalProducts: document.querySelector("#totalProducts"),
  inventoryValue: document.querySelector("#inventoryValue"),
  lowStock: document.querySelector("#lowStock"),
  totalUsers: document.querySelector("#totalUsers"),
};

const exampleProducts = [
  {
    name: "Wireless Noise-Canceling Headphones",
    description: "Over-ear Bluetooth headphones with 40-hour battery life and travel case.",
    price: 129.99,
    stock: 18,
  },
  {
    name: "Smart Fitness Watch",
    description: "Heart-rate tracking, sleep insights, GPS workouts, and water-resistant body.",
    price: 89.5,
    stock: 7,
  },
  {
    name: "Ergonomic Office Chair",
    description: "Adjustable lumbar support, breathable mesh, and smooth rolling casters.",
    price: 214.0,
    stock: 4,
  },
  {
    name: "Ceramic Pour-Over Coffee Set",
    description: "Dripper, glass server, reusable filter, and starter pack of paper filters.",
    price: 46.75,
    stock: 22,
  },
  {
    name: "Portable Power Bank 20000mAh",
    description: "Fast-charging USB-C battery pack for phones, tablets, and travel gear.",
    price: 39.99,
    stock: 0,
  },
  {
    name: "Organic Cotton Hoodie",
    description: "Midweight fleece hoodie with reinforced seams and relaxed everyday fit.",
    price: 58.0,
    stock: 13,
  },
  {
    name: "Stainless Steel Water Bottle",
    description: "Insulated 24 oz bottle that keeps drinks cold for 24 hours.",
    price: 24.95,
    stock: 31,
  },
  {
    name: "Compact Mechanical Keyboard",
    description: "Hot-swappable 75 percent keyboard with tactile switches and white backlight.",
    price: 74.99,
    stock: 5,
  },
  {
    name: "Minimal Desk Lamp",
    description: "Dimmable LED lamp with warm/cool modes and a small charging base.",
    price: 32.25,
    stock: 9,
  },
  {
    name: "Canvas Weekender Bag",
    description: "Durable carry-on bag with shoe pocket, brass hardware, and padded strap.",
    price: 67.4,
    stock: 2,
  },
];

const money = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

function setMessage(element, text, type = "") {
  element.textContent = text;
  element.className = `message ${type}`.trim();
}

function authHeaders() {
  return {
    "Content-Type": "application/json",
    Authorization: `Bearer ${state.token}`,
  };
}

async function request(path, options = {}) {
  const response = await fetch(path, options);
  const text = await response.text();
  const data = text ? JSON.parse(text) : {};

  if (!response.ok) {
    throw new Error(data.error || data.message || "Request failed");
  }

  return data;
}

function normalizeProduct(product) {
  return {
    id: product.ID,
    name: product.Name,
    description: product.Description || "",
    price: Number(product.Price || 0),
    stock: Number(product.Stock || 0),
    createdAt: product.CreatedAt,
  };
}

function normalizeUser(user) {
  return {
    id: user.id || user.ID,
    email: user.email || user.Email,
    createdAt: user.createdAt || user.CreatedAt,
  };
}

function statusFor(stock) {
  if (stock <= 0) return { label: "Out of stock", value: "out" };
  if (stock <= 5) return { label: "Low stock", value: "low" };
  return { label: "Available", value: "available" };
}

function setAuthMode(mode) {
  state.mode = mode;
  els.loginTab.classList.toggle("active", mode === "login");
  els.registerTab.classList.toggle("active", mode === "register");
  els.authSubmit.textContent = mode === "login" ? "Sign in" : "Create account";
  els.password.autocomplete = mode === "login" ? "current-password" : "new-password";
  setMessage(els.authMessage, "");
}

function showApp() {
  const signedIn = Boolean(state.token);
  els.authPanel.classList.toggle("hidden", signedIn);
  els.dashboard.classList.toggle("hidden", !signedIn);
  els.logoutButton.classList.toggle("hidden", !signedIn);
  els.sessionState.textContent = signedIn ? "Signed in" : "Signed out";

  if (signedIn) {
    loadProducts();
    loadUsers();
  }
}

async function login(email, password) {
  const data = await request("/auth/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  state.token = data.token;
  localStorage.setItem("product_manager_token", state.token);
}

async function handleAuth(event) {
  event.preventDefault();
  const email = els.email.value.trim();
  const password = els.password.value;

  setMessage(els.authMessage, state.mode === "login" ? "Signing in..." : "Creating account...");

  try {
    if (state.mode === "register") {
      await request("/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
    }

    await login(email, password);
    els.authForm.reset();
    setMessage(els.authMessage, "");
    showApp();
  } catch (error) {
    setMessage(els.authMessage, error.message, "error");
  }
}

async function loadProducts() {
  try {
    const params = new URLSearchParams({
      page: String(state.productPage),
      limit: String(state.productLimit),
    });
    const query = els.searchProducts.value.trim();
    const stock = els.stockFilter.value;
    if (query) params.set("search", query);
    if (stock !== "all") params.set("stock", stock);

    const data = await request(`/api/products?${params.toString()}`, {
      headers: authHeaders(),
    });
    state.products = (data.products || []).map(normalizeProduct);
    state.productPage = Number(data.page || state.productPage);
    state.productLimit = Number(data.limit || state.productLimit);
    state.productTotal = Number(data.total || 0);
    state.productTotalPages = Number(data.totalPages || 1);

    if (state.products.length === 0 && state.productPage > 1) {
      state.productPage -= 1;
      await loadProducts();
      return;
    }

    renderProducts();
  } catch (error) {
    if (error.message.toLowerCase().includes("token") || error.message.toLowerCase().includes("authorization")) {
      logout();
    }
    setMessage(els.productMessage, error.message, "error");
  }
}

async function loadUsers() {
  try {
    const data = await request("/api/users", {
      headers: authHeaders(),
    });
    state.users = (data.users || []).map(normalizeUser);
    renderUsers();
  } catch (error) {
    if (error.message.toLowerCase().includes("token") || error.message.toLowerCase().includes("authorization")) {
      logout();
    }
  }
}

function filteredProducts() {
  return state.products;
}

function filteredUsers() {
  const query = els.searchUsers.value.trim().toLowerCase();
  return state.users.filter((user) => {
    return [user.id, user.email].join(" ").toLowerCase().includes(query);
  });
}

function renderStats() {
  const value = state.products.reduce((sum, product) => sum + product.price * product.stock, 0);
  const low = state.products.filter((product) => product.stock > 0 && product.stock <= 5).length;

  els.totalProducts.textContent = state.productTotal;
  els.inventoryValue.textContent = money.format(value);
  els.lowStock.textContent = low;
  els.totalUsers.textContent = state.users.length;
  els.navProductCount.textContent = state.productTotal;
  els.navUserCount.textContent = state.users.length;
}

function renderProducts() {
  renderStats();
  const rows = filteredProducts();

  els.productRows.innerHTML = rows.map((product) => {
    const status = statusFor(product.stock);
    const created = product.createdAt ? new Date(product.createdAt).toLocaleDateString() : "-";

    return `
      <tr>
        <td>
          <div class="product-name">${escapeHtml(product.name)}</div>
          <div class="description">${escapeHtml(product.description || "No description")}</div>
        </td>
        <td>${money.format(product.price)}</td>
        <td>${product.stock}</td>
        <td><span class="pill ${status.value}">${status.label}</span></td>
        <td>${created}</td>
        <td>
          <div class="row-actions">
            <button class="ghost" type="button" data-action="edit" data-id="${product.id}">Edit</button>
            <button class="danger" type="button" data-action="delete" data-id="${product.id}">Delete</button>
          </div>
        </td>
      </tr>
    `;
  }).join("");

  els.emptyState.classList.toggle("hidden", rows.length > 0);
  els.productPageInfo.textContent = `Page ${state.productPage} of ${state.productTotalPages} - ${state.productTotal} products`;
  els.prevProductsPage.disabled = state.productPage <= 1;
  els.nextProductsPage.disabled = state.productPage >= state.productTotalPages;
}

function renderUsers() {
  renderStats();
  const rows = filteredUsers();

  els.userRows.innerHTML = rows.map((user) => {
    const created = user.createdAt ? new Date(user.createdAt).toLocaleDateString() : "-";

    return `
      <tr>
        <td>#${user.id}</td>
        <td>
          <div class="product-name">${escapeHtml(user.email)}</div>
          <div class="description">Registered customer account</div>
        </td>
        <td>${created}</td>
      </tr>
    `;
  }).join("");

  els.emptyUsersState.classList.toggle("hidden", rows.length > 0);
}

function setDashboardView(view) {
  state.view = view;
  els.productsViewTab.classList.toggle("active", view === "products");
  els.usersViewTab.classList.toggle("active", view === "users");
  els.productsView.classList.toggle("hidden", view !== "products");
  els.usersView.classList.toggle("hidden", view !== "users");
  els.viewTitle.textContent = view === "products" ? "Products" : "Users";
  els.openProductDrawer.classList.toggle("hidden", view !== "products");
  els.addExamples.classList.toggle("hidden", view !== "products");

  if (view === "users") {
    loadUsers();
    closeProductDrawer();
  }
}

function openProductDrawer() {
  els.productDrawer.classList.remove("hidden");
  els.drawerBackdrop.classList.remove("hidden");
  document.body.classList.add("drawer-open");
  requestAnimationFrame(() => {
    els.productDrawer.classList.add("open");
    els.drawerBackdrop.classList.add("open");
  });
  els.productName.focus();
}

function closeProductDrawer() {
  els.productDrawer.classList.remove("open");
  els.drawerBackdrop.classList.remove("open");
  document.body.classList.remove("drawer-open");
  window.setTimeout(() => {
    if (!els.productDrawer.classList.contains("open")) {
      els.productDrawer.classList.add("hidden");
      els.drawerBackdrop.classList.add("hidden");
    }
  }, 180);
}

function escapeHtml(value) {
  return String(value)
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}

function productPayload() {
  return {
    name: els.productName.value.trim(),
    description: els.productDescription.value.trim(),
    price: Number(els.productPrice.value),
    stock: Number(els.productStock.value),
  };
}

async function saveProduct(event) {
  event.preventDefault();
  const payload = productPayload();
  const editing = state.editingId !== null;

  setMessage(els.productMessage, editing ? "Updating product..." : "Creating product...");

  try {
    await request(editing ? `/api/products/${state.editingId}` : "/api/products", {
      method: editing ? "PUT" : "POST",
      headers: authHeaders(),
      body: JSON.stringify(payload),
    });

    setMessage(els.productMessage, editing ? "Product updated." : "Product created.", "success");
    resetProductForm();
    await loadProducts();
    closeProductDrawer();
  } catch (error) {
    setMessage(els.productMessage, error.message, "error");
  }
}

async function addExampleProducts() {
  const existingNames = new Set(state.products.map((product) => product.name.toLowerCase()));
  const productsToCreate = exampleProducts.filter((product) => !existingNames.has(product.name.toLowerCase()));

  if (productsToCreate.length === 0) {
    setMessage(els.productMessage, "Example products are already in the catalog.", "success");
    return;
  }

  setMessage(els.productMessage, `Adding ${productsToCreate.length} example products...`);
  els.addExamples.disabled = true;

  try {
    await Promise.all(productsToCreate.map((product) => request("/api/products", {
      method: "POST",
      headers: authHeaders(),
      body: JSON.stringify(product),
    })));
    setMessage(els.productMessage, "Example products added.", "success");
    await loadProducts();
  } catch (error) {
    setMessage(els.productMessage, error.message, "error");
  } finally {
    els.addExamples.disabled = false;
  }
}

function editProduct(id) {
  const product = state.products.find((item) => item.id === id);
  if (!product) return;

  state.editingId = id;
  els.formTitle.textContent = "Edit product";
  els.saveProduct.textContent = "Update product";
  els.cancelEdit.classList.remove("hidden");
  els.productName.value = product.name;
  els.productDescription.value = product.description;
  els.productPrice.value = product.price.toFixed(2);
  els.productStock.value = product.stock;
  setMessage(els.productMessage, "");
  openProductDrawer();
}

async function deleteProduct(id) {
  const product = state.products.find((item) => item.id === id);
  const label = product ? `"${product.name}"` : "this product";

  if (!window.confirm(`Delete ${label}? This cannot be undone.`)) {
    return;
  }

  try {
    await request(`/api/products/${id}`, {
      method: "DELETE",
      headers: authHeaders(),
    });
    setMessage(els.productMessage, "Product deleted.", "success");
    await loadProducts();
  } catch (error) {
    setMessage(els.productMessage, error.message, "error");
  }
}

function resetProductForm() {
  state.editingId = null;
  els.productForm.reset();
  els.formTitle.textContent = "Add product";
  els.saveProduct.textContent = "Create product";
}

function startCreateProduct() {
  resetProductForm();
  setMessage(els.productMessage, "");
  openProductDrawer();
}

function logout() {
  state.token = "";
  state.products = [];
  state.users = [];
  localStorage.removeItem("product_manager_token");
  resetProductForm();
  closeProductDrawer();
  showApp();
}

els.loginTab.addEventListener("click", () => setAuthMode("login"));
els.registerTab.addEventListener("click", () => setAuthMode("register"));
els.productsViewTab.addEventListener("click", () => setDashboardView("products"));
els.usersViewTab.addEventListener("click", () => setDashboardView("users"));
els.authForm.addEventListener("submit", handleAuth);
els.logoutButton.addEventListener("click", logout);
els.productForm.addEventListener("submit", saveProduct);
els.openProductDrawer.addEventListener("click", startCreateProduct);
els.closeProductDrawer.addEventListener("click", () => {
  resetProductForm();
  setMessage(els.productMessage, "");
  closeProductDrawer();
});
els.drawerBackdrop.addEventListener("click", () => {
  resetProductForm();
  setMessage(els.productMessage, "");
  closeProductDrawer();
});
els.addExamples.addEventListener("click", addExampleProducts);
els.cancelEdit.addEventListener("click", () => {
  resetProductForm();
  setMessage(els.productMessage, "");
  closeProductDrawer();
});
els.searchProducts.addEventListener("input", () => {
  state.productPage = 1;
  loadProducts();
});
els.stockFilter.addEventListener("change", () => {
  state.productPage = 1;
  loadProducts();
});
els.prevProductsPage.addEventListener("click", () => {
  if (state.productPage <= 1) return;
  state.productPage -= 1;
  loadProducts();
});
els.nextProductsPage.addEventListener("click", () => {
  if (state.productPage >= state.productTotalPages) return;
  state.productPage += 1;
  loadProducts();
});
els.searchUsers.addEventListener("input", renderUsers);
els.productRows.addEventListener("click", (event) => {
  const button = event.target.closest("button[data-action]");
  if (!button) return;

  const id = Number(button.dataset.id);
  if (button.dataset.action === "edit") {
    editProduct(id);
  }
  if (button.dataset.action === "delete") {
    deleteProduct(id);
  }
});

setAuthMode("login");
showApp();
