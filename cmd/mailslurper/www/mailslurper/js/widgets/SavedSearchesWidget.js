// Copyright 2013-2016 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery",
		"services/SettingsService",
		"moment",

		"hbs!templates/savedSearchesModal",
		"hbs!templates/saveSearchModal",

		"bootstrap-dialog"
	],
	function($, settingsService, moment, savedSearchesModalTemplate, saveSearchModalTemplate, Dialog) {
		"use strict";

		var widget = {
			showPicker: function(callback) {
				var dialogRef = Dialog.show({
					title: "Saved Searches",
					message: savedSearchesModalTemplate({ savedSearches: settingsService.retrieveSavedSearches() }),
					closable: true,
					nl2br: false,
					buttons: [
						{
							id: "btnManageSavedSearches",
							label: "Manage",
							cssClass: "btn-default",
							action: function() {
								window.location = "/savedsearches";
							}
						},
						{
							id: "btnCancelSavedSearch",
							label: "Cancel",
							cssClass: "btn-default",
							action: function(dialogRef) {
								dialogRef.close();
							}
						},
						{
							id: "btnOK",
							label: "OK",
							cssClass: "btn-primary",
							action: function() {
								var savedSearch = settingsService.getSavedSearchByIndex(window.parseInt($("#savedSearchID option:selected").val(), 10));
								callback(savedSearch);
								dialogRef.close();
							}
						}
					]
				});
			},

			showSaveSearchModal: function(callback) {
				var dialogRef = Dialog.show({
					title: "Save Search",
					message: saveSearchModalTemplate(),
					closable: true,
					nl2br: false,
					buttons: [
						{
							id: "btnCancelSaveSearch",
							label: "Cancel",
							cssClass: "btn-default",
							action: function(dialogRef) {
								dialogRef.close();
							}
						},
						{
							id: "btnSaveSearchOK",
							label: "OK",
							cssClass: "btn-primary",
							action: function(dialogRef) {
								var saveSearchName = $("#txtSaveSearchName").val();
								if (saveSearchName.length <= 0) {
									alert("Please enter a name for your search!");
								} else {
									dialogRef.close();
									callback(saveSearchName);
								}
							}
						}
					],
					onshown: function(dialogRef) {
						$("#txtSaveSearchName").focus();
					}
				});
			}
		};

		return widget;
	}
);
