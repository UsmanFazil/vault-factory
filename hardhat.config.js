require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-etherscan");
// require('@openzeppelin/hardhat-upgrades');
// require("hardhat-gas-reporter");
// require('solidity-coverage');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.17",
  defaultNetwork: "hardhat",
  solidity: {
    compilers: [
      {
        version: "0.8.17",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          },
        }
      },
      {
        version: "0.6.12",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      },
    ]
  },
  // gasReporter: {
  //   currency: 'USD',
  //   gasPrice: 120,
  //   enabled:  true, //(process.env.REPORT_GAS)? true :
  //   ethPrice: 1148.1
  //   // coinmarketcap: "8cab7a0f-7baf-4e90-8c12-39d78eacb364"//process.env.COINMARKETCAP_API_KEY
  // },
  settings: {
    optimizer: {
      enabled: true,
      runs: 200,
    },
  },
  networks: {
    
    hardhat: {
      allowUnlimitedContractSize: true,
      timeout: 1200000000000000
    },

  },
  etherscan: {
    // Your API key for Etherscan
    // Obtain one at https://etherscan.io/
    apiKey: "FG6FG5FQJNTWMD6IUIC5WITIZC8C2ECTHQ"
  }

};
