function getAllCategories() {
  const vm = this;
  return fetch('/api/categories')
  .then(data => data.json())
  .then(data => data.data)
  .then(data => {
    vm.categories = data;
  });
};
