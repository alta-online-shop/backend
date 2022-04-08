function getAllCategories() {
  const vm = this;
  return fetch('/api/categories')
  .then(data => data.json())
  .then(data => data.data)
  .then(data => {
    vm.categories = data;
  });
};

function createCategory(e) {
  e.preventDefault();
  const vm = this;
  const body = JSON.stringify(this.category);
  console.log(body);
  return fetch('/api/categories', {
    method: 'POST',
    body,
  })
  .then(() => {
    if (!vm.addMore) window.history.back(-1);
    vm.category = {
      name: '',
      description: '',
    };
    vm.addMore = false;
  })
  .catch(console.log);
}
