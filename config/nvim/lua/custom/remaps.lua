local telescope = require('telescope.builtin')
vim.keymap.set('n', '<leader>ff', telescope.find_files, { desc = 'Telescope find files' })
vim.keymap.set('n', '<leader>gf', telescope.git_files, { desc = 'Telescope git files' })
vim.keymap.set("n", "<leader>e", function()
  vim.diagnostic.open_float()
end, { desc = "Show diagnostics in a floating window" })

