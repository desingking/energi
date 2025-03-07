// Copyright 2018 The Energi Core Authors
// Copyright 2015 The go-ethereum Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

// package web3ext contains geth specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"accounting": Accounting_JS,
	"admin":      Admin_JS,
	//"chequebook": Chequebook_JS,
	//"clique":     Clique_JS,
	"energi": Energi_JS,
	//"ethash":     Ethash_JS,
	"debug":      Debug_JS,
	"eth":        Eth_JS,
	"masternode": Masternode_JS,
	"miner":      Miner_JS,
	"net":        Net_JS,
	"personal":   Personal_JS,
	"rpc":        RPC_JS,
	"shh":        Shh_JS,
	"swarmfs":    SWARMFS_JS,
	"txpool":     TxPool_JS,
}

const Chequebook_JS = `
web3._extend({
	property: 'chequebook',
	methods: [
		new web3._extend.Method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const Clique_JS = `
web3._extend({
	property: 'clique',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.Method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const Energi_JS = `
web3._extend.formatters.outputProposalFormatter = function(item){
	var toDecimal = web3._extend.utils.toDecimal;
	return {
		proposal:     item.Proposal,
		proposer:     item.Proposer,
		createdBlock: item.CreatedBlock,
		deadline:     item.Deadline,
		quorumWeight: toDecimal(item.QuorumWeight),
		totalWeight:  toDecimal(item.TotalWeight),
		rejectWeight: toDecimal(item.RejectWeight),
		acceptWeight: toDecimal(item.AcceptWeight),
		finished:     item.Finished,
		accepted:     item.Accepted,
		balance:      toDecimal(item.Balance),
	};
};

web3._extend.formatters.coinSearchFormatter = function(list){
	var toDecimal = web3._extend.utils.toDecimal;
	for (var i = 0; i < list.length; ++i) {
		var item = list[i];
		item.Amount = toDecimal(item.Amount);
	}
	return list;
};

web3._extend({
	property: 'energi',
	methods: [
		// Migration
		new web3._extend.Method({
			name: 'listGen2Coins',
			call: 'energi_listGen2Coins',
			params: 0,
			outputFormatter: web3._extend.formatters.coinSearchFormatter,
		}),
		new web3._extend.Method({
			name: 'searchGen2Coins',
			call: 'energi_searchGen2Coins',
			params: 2,
			outputFormatter: web3._extend.formatters.coinSearchFormatter,
		}),
		new web3._extend.Method({
			name: 'searchGen3DestinationByGen2Address',
			call: 'energi_searchGen3DestinationByGen2Address',
			params: 2,
			outputFormatter: web3._extend.formatters.coinSearchFormatter,
		}),
		new web3._extend.Method({
			name: 'searchRawGen2Coins',
			call: 'energi_searchRawGen2Coins',
			params: 2,
			outputFormatter: web3._extend.formatters.coinSearchFormatter,
		}),
		new web3._extend.Method({
			name: 'claimGen2CoinsDirect',
			call: 'energi_claimGen2CoinsDirect',
			params: 3,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'claimGen2CoinsCombined',
			call: 'energi_claimGen2CoinsCombined',
			params: 3
		}),
		new web3._extend.Method({
			name: 'claimGen2CoinsImport',
			call: 'energi_claimGen2CoinsImport',
			params: 2
		}),

		// Blacklist
		new web3._extend.Method({
			name: 'blacklistInfo',
			call: 'energi_blacklistInfo',
			params: 0
			outputFormatter: function(list) {
				var res = [];
				var proposalf = web3._extend.formatters.outputProposalFormatter;
				for (var i = 0; i < list.length; ++i) {
					var item = list[i];
					res.push({
						target:  item.Target,
						enforce: proposalf(item.Enforce),
						revoke:  item.Revoke && proposalf(item.Revoke),
						drain:   item.Drain && proposalf(item.Drain),
						blocked: item.Blocked,
					});
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'blacklistEnforce',
			call: 'energi_blacklistEnforce',
			params: 4
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'blacklistRevoke',
			call: 'energi_blacklistRevoke',
			params: 4
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'blacklistDrain',
			call: 'energi_blacklistDrain',
			params: 4
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'blacklistCollect',
			call: 'energi_blacklistCollect',
			params: 3
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),

		// Governance
		new web3._extend.Method({
			name: 'voteAccept',
			call: 'energi_voteAccept',
			params: 3
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'voteReject',
			call: 'energi_voteReject',
			params: 3
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'withdrawFee',
			call: 'energi_withdrawFee',
			params: 3
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),

		// Governance upgrades
		new web3._extend.Method({
			name: 'upgradeInfo',
			call: 'energi_upgradeInfo',
			params: 0
			outputFormatter: function(status) {
				var res = {};
				var proposalf = web3._extend.formatters.outputProposalFormatter;
				var keys = Object.keys(status);
				for (var i = 0; i < keys.length; ++i) {
					var k = keys[i];
					var items = status[k];
					var res_items = res[k] = [];

					for (var j = 0; j < items.length; ++j) {
						var item = items[j];
						var res_item = proposalf(item);
						res_item.impl = item.Impl;
						res_item.proxy = item.Proxy;
						res_items.push(res_item);
					}
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'upgradePropose',
			call: 'energi_upgradePropose',
			params: 6
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'upgradePerform',
			call: 'energi_upgradePerform',
			params: 4
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'upgradeCollect',
			call: 'energi_upgradeCollect',
			params: 4
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),

		// Governance budget
		new web3._extend.Method({
			name: 'budgetInfo',
			call: 'energi_budgetInfo',
			params: 0
			outputFormatter: function(status) {
				var proposals = [];
				var toDecimal = web3._extend.utils.toDecimal;
				var res = {
					balance: toDecimal(status.Balance),
					proposals: proposals,
				};
				var proposalf = web3._extend.formatters.outputProposalFormatter;
				var raw_proposals = status.Proposals;
				for (var i = 0; i < raw_proposals.length; ++i) {
					var raw_item = raw_proposals[i];
					var item = proposalf(raw_item);
					item.proposedAmount = toDecimal(raw_item.ProposedAmount);
					item.paidAmount = toDecimal(raw_item.PaidAmount);
					item.refUUID = raw_item.RefUUID;
					proposals.push(item);
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'budgetPropose',
			call: 'energi_budgetPropose',
			params: 6
			inputFormatter: [
				web3._extend.utils.fromDecimal,
				null,
				null,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),


		// Compensation Fund
		new web3._extend.Method({
			name: 'compensationInfo',
			call: 'energi_compensationInfo',
			params: 0
			outputFormatter: function(status) {
				var proposals = [];
				var toDecimal = web3._extend.utils.toDecimal;
				var res = {
					balance: toDecimal(status.Balance),
					proposals: proposals,
				};
				var proposalf = web3._extend.formatters.outputProposalFormatter;
				var raw_proposals = status.Proposals;
				for (var i = 0; i < raw_proposals.length; ++i) {
					var raw_item = raw_proposals[i];
					var item = proposalf(raw_item);
					item.proposedAmount = toDecimal(raw_item.ProposedAmount);
					item.paidAmount = toDecimal(raw_item.PaidAmount);
					item.refUUID = raw_item.RefUUID;
					proposals.push(item);
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'compensationPropose',
			call: 'energi_compensationPropose',
			params: 6
			inputFormatter: [
				web3._extend.utils.fromDecimal,
				null,
				null,
				web3._extend.utils.fromDecimal,
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'compensationProcess',
			call: 'energi_compensationProcess',
			params: 2
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				null,
			],
			outputFormatter: console.log,
		}),

		// Checkpoints
		new web3._extend.Method({
			name: 'checkpointInfo',
			call: 'energi_checkpointInfo',
			params: 0
			outputFormatter: function(status) {
				var res = {
					registry: [],
					active: [],
				};
				var cpinfo = function(cp) {
					return {
						number: cp.Number,
						hash: cp.Hash,
						since: cp.Since,
						sigCount: cp.SigCount,
					};
				};
				var raw_registry = status.Registry;
				for (var i = 0; i < raw_registry.length; ++i) {
					var raw_item = raw_registry[i];
					res.registry.push(cpinfo(raw_registry[i]));
				}
				var raw_active = status.Active;
				for (var i = 0; i < raw_active.length; ++i) {
					var raw_item = raw_active[i];
					res.active.push(cpinfo(raw_active[i]));
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'checkpointPropose',
			call: 'energi_checkpointPropose',
			params: 3
			inputFormatter: [
				null,
				null,
				null,
			],
			outputFormatter: console.log,
		}),

		// hardfork registry
		new web3._extend.Method({
			name: 'hardforkList'
			call: 'energi_hardforkEnumerate',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'hardforkListActive'
			call: 'energi_hardforkEnumerateActive',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'hardforkListPending'
			call: 'energi_hardforkEnumeratePending',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'hardforkInfo'
			call: 'energi_hardforkGet',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'hardforkIsActive'
			call: 'energi_hardforkIsActive',
			params: 1,
		})
	],
	properties: [
	]
});
`

const Ethash_JS = `
web3._extend({
	property: 'ethash',
	methods: [
		new web3._extend.Method({
			name: 'getWork',
			call: 'ethash_getWork',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'ethash_getHashrate',
			params: 0
		}),
		new web3._extend.Method({
			name: 'submitWork',
			call: 'ethash_submitWork',
			params: 3,
		}),
		new web3._extend.Method({
			name: 'submitHashRate',
			call: 'ethash_submitHashRate',
			params: 2,
		}),
	]
});
`

const Admin_JS = `
web3._extend({
	property: 'admin',
	methods: [
		new web3._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'addTrustedPeer',
			call: 'admin_addTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removeTrustedPeer',
			call: 'admin_removeTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.Method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
		new web3._extend.Method({
			name: 'checkpointLocal',
			call: 'admin_checkpointLocal',
			params: 2,
			inputFormatter: [
				null,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'validateMigration',
			call: 'admin_validateMigration',
			params: 1,
			outputFormatter: console.log,
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const Debug_JS = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.Method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.Method({
			name: 'metrics',
			call: 'debug_metrics',
			params: 1
		}),
		new web3._extend.Method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.Method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.Method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.Method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.Method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'mutexProfile',
			call: 'debug_mutexProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setMutexProfileFraction',
			call: 'debug_setMutexProfileFraction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMutexProfile',
			call: 'debug_writeMutexProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBadBlock',
			call: 'debug_traceBadBlock',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBadBlockToFile',
			call: 'debug_standardTraceBadBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBlockToFile',
			call: 'debug_standardTraceBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
	],
	properties: []
});
`

const Eth_JS = `
web3._extend({
	property: 'eth',
	methods: [
		new web3._extend.Method({
			name: 'chainId',
			call: 'eth_chainId',
			params: 0
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'eth_sign',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'resend',
			call: 'eth_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'eth_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'submitTransaction',
			call: 'eth_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'getRawTransaction',
			call: 'eth_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'eth_getRawTransactionByBlockHashAndIndex' : 'eth_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
		new web3._extend.Method({
			name: 'getProof',
			call: 'eth_getProof',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null, web3._extend.formatters.inputBlockNumberFormatter]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'eth_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(web3._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`

const Masternode_JS = `
web3._extend.formatters.outputMasternodeFormatter = function(item){
	return {
		masternode:     item.Masternode,
		owner:          item.Owner,
		enode:          item.Enode,
		collateral:     web3._extend.utils.toDecimal(item.Collateral),
		announcedBlock: item.AnnouncedBlock,
		isActive:       item.IsActive,
		isAlive:        item.IsAlive,
		swFeatures:     item.SWFeatures,
		swVersion:      item.SWVersion,
	};
};

web3._extend({
	property: 'masternode',
	methods: [
		new web3._extend.Method({
			name: 'collateralBalance',
			call: 'masternode_collateralBalance',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter],
			outputFormatter: function(status) {
				return {
					balance: web3._extend.utils.toDecimal(status.Balance),
					lastBlock: web3._extend.utils.toDecimal(status.LastBlock),
				};
			}
		}),
		new web3._extend.Method({
			name: 'depositCollateral',
			call: 'masternode_depositCollateral',
			params: 3,
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.utils.fromDecimal,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'withdrawCollateral',
			call: 'masternode_withdrawCollateral',
			params: 3,
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.utils.fromDecimal,
				null,
			],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'listMasternodes',
			call: 'masternode_listMasternodes',
			params: 0,
			outputFormatter: function(list) {
				var res = [];
				for (var i = 0; i < list.length; ++i) {
					res.push(web3._extend.formatters.outputMasternodeFormatter(list[i]));
				}
				return res;
			},
		}),
		new web3._extend.Method({
			name: 'masternodeInfo',
			call: 'masternode_masternodeInfo',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter],
			outputFormatter: function(status) {
				return web3._extend.formatters.outputMasternodeFormatter(status);
			}
		}),
		new web3._extend.Method({
			name: 'stats',
			call: 'masternode_stats',
			params: 0,
			outputFormatter: function(status) {
				return {
					active: status.Active,
					total: status.Total,
					activeCollateral: web3._extend.utils.toDecimal(status.ActiveCollateral),
					totalCollateral: web3._extend.utils.toDecimal(status.TotalCollateral),
					maxOfAllTimes: web3._extend.utils.toDecimal(status.MaxOfAllTimes),
				};
			}
		}),
		new web3._extend.Method({
			name: 'announce',
			call: 'masternode_announce',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null, null],
		}),
		new web3._extend.Method({
			name: 'denounce',
			call: 'masternode_denounce',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null],
		}),
	],
	properties: []
});
`

const Miner_JS = `
web3._extend({
	property: 'miner',
	methods: [
		new web3._extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'stop',
			call: 'miner_stop'
		}),
		new web3._extend.Method({
			name: 'setEtherbase',
			call: 'miner_setEtherbase',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'setRecommitInterval',
			call: 'miner_setRecommitInterval',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
		new web3._extend.Method({
			name: 'addDPoS',
			call: 'miner_addDPoS',
			params: 2,
			inputFormatter: [
				web3._extend.formatters.inputAddressFormatter,
				web3._extend.formatters.inputAddressFormatter,
			]
		}),
		new web3._extend.Method({
			name: 'removeDPoS',
			call: 'miner_removeDPoS',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'setNonceCap',
			call: 'miner_setNonceCap',
			params: 1,
			inputFormatter: [null],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'setAutocollateralize',
			call: 'miner_setAutocollateralize',
			params: 1,
			inputFormatter: [null],
			outputFormatter: console.log,
		}),
		new web3._extend.Method({
			name: 'stakingStatus',
			call: 'miner_stakingStatus',
			params: 0
			outputFormatter: function(status) {
				var res = {
					hash: status.Hash,
					height: status.Height,
					miner: status.Miner,
					nonceCap: status.NonceCap,
					staking: status.Staking,
					totalWeight: status.TotalWeight,
					accounts: [],
				};
				var raw_accounts = status.Accounts;
				for (var i = 0; i < raw_accounts.length; ++i) {
					var raw_acct = raw_accounts[i];
					res.accounts.push({
						account: raw_accounts[i].Account,
						weight: raw_accounts[i].Weight,
					});
				}
				return res;
			},
		}),
	],
	properties: []
});
`

const Net_JS = `
web3._extend({
	property: 'net',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const Personal_JS = `
web3._extend({
	property: 'personal',
	methods: [
		new web3._extend.Method({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new web3._extend.Method({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Method({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new web3._extend.Method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.Method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'personal_signTransaction',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, null]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RPC_JS = `
web3._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const Shh_JS = `
web3._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SWARMFS_JS = `
web3._extend({
	property: 'swarmfs',
	methods:
	[
		new web3._extend.Method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.Method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.Method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxPool_JS = `
web3._extend({
	property: 'txpool',
	methods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = web3._extend.utils.toDecimal(status.pending);
				status.queued = web3._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`

const Accounting_JS = `
web3._extend({
	property: 'accounting',
	methods: [
		new web3._extend.Property({
			name: 'balance',
			getter: 'account_balance'
		}),
		new web3._extend.Property({
			name: 'balanceCredit',
			getter: 'account_balanceCredit'
		}),
		new web3._extend.Property({
			name: 'balanceDebit',
			getter: 'account_balanceDebit'
		}),
		new web3._extend.Property({
			name: 'bytesCredit',
			getter: 'account_bytesCredit'
		}),
		new web3._extend.Property({
			name: 'bytesDebit',
			getter: 'account_bytesDebit'
		}),
		new web3._extend.Property({
			name: 'msgCredit',
			getter: 'account_msgCredit'
		}),
		new web3._extend.Property({
			name: 'msgDebit',
			getter: 'account_msgDebit'
		}),
		new web3._extend.Property({
			name: 'peerDrops',
			getter: 'account_peerDrops'
		}),
		new web3._extend.Property({
			name: 'selfDrops',
			getter: 'account_selfDrops'
		}),
	]
});
`
