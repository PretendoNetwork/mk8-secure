package main

/*func getAssociatedNexUniqueIdWithMyPrincipalId(err error, client *nex.Client, callID uint32) {
	rmcResponseStream := nex.NewStreamOut(nexServer)

	rmcResponseStream.WriteUInt64LE(0)
	rmcResponseStream.WriteUInt64LE(0)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(nexproto.UtilityProtocolID, callID)
	rmcResponse.SetSuccess(nexproto.UtilityMethodGetAssociatedNexUniqueIdWithMyPrincipalId, rmcResponseBody)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	nexServer.Send(responsePacket)
}
*/
