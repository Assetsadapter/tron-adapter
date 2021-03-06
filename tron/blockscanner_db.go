/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tron

import (
	"fmt"
	"github.com/blocktree/openwallet/v2/openwallet"
)

// //GetCurrentBlockHeader 获取当前区块高度
// func (bs *TronBlockScanner) GetCurrentBlockHeader() (*openwallet.BlockHeader, error) {

// 	var (
// 		block       *core.Block
// 		blockHeight uint64 = 0
// 		hash        string
// 		err         error
// 	)

// 	blockHeight, hash = bs.GetLocalNewBlock()

// 	//如果本地没有记录，查询接口的高度
// 	if blockHeight == 0 {
// 		block, hash, err = bs.wm.GetNowBlock()
// 		if err != nil {
// 			return nil, err
// 		}
// 		blockHeight = uint64(block.GetBlockHeader().GetRawData().GetNumber())
// 	}

// 	return &openwallet.BlockHeader{Height: blockHeight, Hash: hash}, nil
// }

//GetLocalNewBlock 获取本地记录的区块高度和hash
func (bs *TronBlockScanner) GetLocalNewBlock() (uint64, string, error) {

	//var (
	//	blockHeight uint64
	//	blockHash   string
	//)
	//
	////获取本地区块高度
	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return 0, ""
	//}
	//defer db.Close()
	//
	//db.Get(blockchainBucket, "blockHeight", &blockHeight)
	//db.Get(blockchainBucket, "blockHash", &blockHash)
	//
	//return blockHeight, blockHash

	if bs.BlockchainDAI == nil {
		return 0, "", fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetCurrentBlockHead(bs.wm.Symbol())
	if err != nil {
		return 0, "", err
	}

	return header.Height, header.Hash, nil
}

//GetLocalBlock 获取本地区块数据
func (bs *TronBlockScanner) GetLocalBlock(height uint64) (*Block, error) {

	//var (
	//	block Block
	//)
	//
	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return nil, err
	//}
	//defer db.Close()
	//
	//err = db.One("Height", height, &block)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &block, nil

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetLocalBlockHeadByHeight(height, bs.wm.Symbol())
	if err != nil {
		return nil, err
	}

	block := &Block{
		Hash:   header.Hash,
		Height: header.Height,
	}

	return block, nil
}

//SaveUnscanRecord 保存交易记录到钱包数据库
func (bs *TronBlockScanner) SaveUnscanRecord(record *openwallet.UnscanRecord) error {

	//if record == nil {
	//	return errors.New("the unscan record to save is nil")
	//}
	//
	////if record.BlockHeight == 0 {
	////	return errors.New("unconfirmed transaction do not rescan")
	////}
	//
	////获取本地区块高度
	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return err
	//}
	//defer db.Close()
	//
	//return db.Save(record)

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.SaveUnscanRecord(record)
}

//SaveLocalNewBlock 记录区块高度和hash到本地
func (bs *TronBlockScanner) SaveLocalNewBlock(blockHeight uint64, blockHash string) error {

	//获取本地区块高度
	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return
	//}
	//defer db.Close()
	//
	//db.Set(blockchainBucket, "blockHeight", &blockHeight)
	//db.Set(blockchainBucket, "blockHash", &blockHash)
	//
	//if bs.BlockchainDAI == nil {
	//	return fmt.Errorf("Blockchain DAI is not setup ")
	//}

	header := &openwallet.BlockHeader{
		Hash:   blockHash,
		Height: blockHeight,
		Fork:   false,
		Symbol: bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveCurrentBlockHead(header)
}

//SaveLocalBlock 记录本地新区块
func (bs *TronBlockScanner) SaveLocalBlock(block *Block) error {

	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return
	//}
	//defer db.Close()
	//
	//db.Save(block)

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	header := &openwallet.BlockHeader{
		Hash:              block.Hash,
		Merkleroot:        "",
		Previousblockhash: block.Previousblockhash,
		Height:            block.Height,
		Time:              uint64(block.Time),
		Symbol:            bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveLocalBlockHead(header)
}

//DeleteUnscanRecord 删除指定高度的未扫记录
func (bs *TronBlockScanner) DeleteUnscanRecord(height uint64) error {
	////获取本地区块高度
	//db, err := storm.Open(filepath.Join(bs.wm.Config.dbPath, bs.wm.Config.BlockchainFile))
	//if err != nil {
	//	return err
	//}
	//defer db.Close()
	//
	//var list []*UnscanRecord
	//err = db.Find("BlockHeight", height, &list)
	//if err != nil {
	//	return err
	//}
	//
	//for _, r := range list {
	//	db.DeleteStruct(r)
	//}
	//
	//return nil

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.DeleteUnscanRecordByHeight(height, bs.wm.Symbol())
}

func (bs *TronBlockScanner) GetUnscanRecords() ([]*openwallet.UnscanRecord, error) {

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.GetUnscanRecords(bs.wm.Symbol())
}
