function getAllProducts() {
  const vm = this;
  return fetch('/api/products')
  .then(data => data.json())
  .then(data => data.data)
  .then(data => {
    vm.products = data;
  });
};

function createProduct(e) {
  e.preventDefault();
  const vm = this;
  const body = JSON.stringify(this.product);
  console.log(body);
  return fetch('/api/products', {
    method: 'POST',
    body,
  })
  .then(() => {
    if (!vm.addMore) window.location.href = '/';
    vm.product = {
      name: '',
      price: 0,
      description: '',
      categories: [],
    };
    vm.addMore = false;
  })
  .catch(console.log);
}

function deleteProduct(id) {
  const vm = this;
  return fetch(`/api/products/${id}`, {
    method: 'DELETE',
  })
  .then(() => {
    vm.getAllProducts();
  })
  .catch(console.log);
}

