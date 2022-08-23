// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";

contract Vote is Ownable {
    mapping(string => string) private votes;
    mapping(string => uint32) private voteCount;
    string[] private options;

    constructor(string[] memory _options) {
        options = _options;
    }

    function castVote(string memory voterCode, string memory option) public onlyOwner {
        require(_isValidOption(option), "Vote: Value sent is not an Option");
        require(!_hasVoted(voterCode), "Vote: Voter has already voted");
        votes[voterCode] = option;
        voteCount[option]++;
    }

    function _isValidOption(string memory option) internal view returns(bool){
        for(uint8 i = 0; i < options.length; i++) {
            if(keccak256(bytes(options[i])) == keccak256(bytes(option))){
                return true;
            }
        }
        return false;
    }

    function _hasVoted(string memory voterCode) internal view returns(bool) {
        return bytes(votes[voterCode]).length > 0;
    }

    function getVote(string memory voterCode) public view onlyOwner returns(string memory){
        require(_hasVoted(voterCode), "Vote: VoterCode not found");
        return votes[voterCode];
    }

    function getVoteCount(string memory option) public view onlyOwner returns(uint32){
        require(_isValidOption(option), "Vote: Value sent is not an Option");
        return voteCount[option];
    }

}