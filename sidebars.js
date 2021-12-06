const fs = require('fs');
const path = require('path');

const rules = fs.readdirSync('./fisherman-source/docs/configuration/rules');

const rulesItems = rules.map(mdFile => ({
  type: "doc",
  id: `configuration/rules/${path.parse(mdFile).name}`
}));

module.exports = {
  someSidebar: [
    {
      type: "doc",
      id: "introduction"
    },
    {
      type: "doc",
      id: "getting-started",
    },
    {
      type: "category",
      label: "Configuration",
      collapsed: false,
      items: [
        {
          type: "doc",
          id: "configuration/configuration-files"
        },
        {
          type: "doc",
          id: "configuration/hooks-configuration"
        },
        {
          type: "category",
          label: "Rules",
          collapsed: true,
          items: rulesItems
        },
        {
          type: "doc",
          id: "configuration/variables"
        },
        {
          type: "doc",
          id: "configuration/expressions"
        },
        {
          type: "doc",
          id: "configuration/output"
        },
      ],
    },
    {
      type: "doc",
      id: "cli"
    },
    {
      type: "doc",
      id: "faq"
    },
  ],
}
