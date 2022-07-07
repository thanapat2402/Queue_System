module.exports = {
  moduleNameMapper: {
    "@core/(.*)": "<rootDir>/src/app/core/$1",
    "@oshc(.*)": "<rootDir>/src/app/modules/oshc/$1",
    "@shared(.*)": "<rootDir>/src/app/@shared/$1",
  },
  preset: "jest-preset-angular",
  setupFilesAfterEnv: ["<rootDir>/setup-jest.ts"],
  moduleFileExtensions: ["json", "js", "jsx", "ts", "tsx", "vue", "cjs"],
  moduleDirectories: ["node_modules", "src"],
  extensionsToTreatAsEsm: [".ts"],
  globals: {
    "ts-jest": {
      useESM: true,
    },
  },
};
