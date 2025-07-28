local lspconfig = require('lspconfig')

lspconfig.gopls.setup({
  on_attach = function(client, bufnr)
    -- Organize imports before saving
    vim.api.nvim_create_autocmd("BufWritePre", {
      buffer = bufnr,
      callback = function()
        local params = vim.lsp.util.make_range_params()
        params.context = { only = { "source.organizeImports" } }
        local result = vim.lsp.buf_request_sync(0, "textDocument/codeAction", params, 1000)
        for _, res in pairs(result or {}) do
          for _, action in pairs(res.result or {}) do
            if action.edit then
              vim.lsp.util.apply_workspace_edit(action.edit, "utf-16")
            else
              vim.lsp.buf.execute_command(action.command)
            end
          end
        end
      end,
    })
  end,
})

-- Format and organize imports
vim.api.nvim_create_autocmd("BufWritePre", {
  buffer = bufnr,
  callback = function()
    vim.lsp.buf.format({ async = false })  -- format code
    -- organize imports (as above)
  end,
})


vim.lsp.enable('gopls')

vim.diagnostic.config({
  virtual_text = true,  -- shows inline error/warning text
  signs = true,         -- shows icons in gutter
  underline = true,     -- underlines problems
  update_in_insert = false, -- don’t show diagnostics while typing
  severity_sort = true,     -- sorts diagnostics by severity
})
