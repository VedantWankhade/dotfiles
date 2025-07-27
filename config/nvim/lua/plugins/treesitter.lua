return {
  'nvim-treesitter/nvim-treesitter',
  build = ':TSUpdate',
  config = function()
    require('nvim-treesitter.configs').setup {
      ensure_installed = { "go", "lua", "javascript", "typescript" }, -- add your langs
      highlight = {
        enable = true,
      },
      indent = {
        enable = true
      },
    }
  end
}

