// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An address book with attributes.
type AddressBook struct {
	// The description of the address book.
	Description *string
	// The ARN of the address book.
	AddressBookArn *string
	// The name of the address book.
	Name *string
}

// Information related to an address book.
type AddressBookData struct {
	// The description of the address book.
	Description *string
	// The name of the address book.
	Name *string
	// The ARN of the address book.
	AddressBookArn *string
}

// The audio message. There is a 1 MB limit on the audio file input and the only
// supported format is MP3. To convert your MP3 audio files to an Alexa-friendly,
// required codec version (MPEG version 2) and bit rate (48 kbps), you might use
// converter software. One option for this is a command-line tool, FFmpeg. For more
// information, see FFmpeg (https://www.ffmpeg.org/). The following command
// converts the provided to an MP3 file that is played in the announcement: ffmpeg
// -i -ac 2 -codec:a libmp3lame -b:a 48k -ar 16000
type Audio struct {
	// The locale of the audio message. Currently, en-US is supported.
	Locale Locale
	// The location of the audio file. Currently, S3 URLs are supported. Only S3
	// locations comprised of safe characters are valid. For more information, see Safe
	// Characters
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#Safe%20Characters).
	Location *string
}

// Usage report with specified parameters.
type BusinessReport struct {
	// The failure code.
	FailureCode BusinessReportFailureCode
	// The S3 location of the output reports.
	S3Location *BusinessReportS3Location
	// The time of report delivery.
	DeliveryTime *time.Time
	// The download link where a user can download the report.
	DownloadUrl *string
	// The status of the report generation execution (RUNNING, SUCCEEDED, or FAILED).
	Status BusinessReportStatus
}

// The content range of the report.
type BusinessReportContentRange struct {
	// The interval of the content range.
	Interval BusinessReportInterval
}

// The recurrence of the reports.
type BusinessReportRecurrence struct {
	// The start date.
	StartDate *string
}

// The S3 location of the output reports.
type BusinessReportS3Location struct {
	// The S3 bucket name of the output reports.
	BucketName *string
	// The path of the business report.
	Path *string
}

// The schedule of the usage report.
type BusinessReportSchedule struct {
	// The name identifier of the schedule.
	ScheduleName *string
	// The content range of the reports.
	ContentRange *BusinessReportContentRange
	// The S3 bucket name of the output reports.
	S3BucketName *string
	// The recurrence of the reports.
	Recurrence *BusinessReportRecurrence
	// The details of the last business report delivery for a specified time interval.
	LastBusinessReport *BusinessReport
	// The S3 key where the report is delivered.
	S3KeyPrefix *string
	// The ARN of the business report schedule.
	ScheduleArn *string
	// The format of the generated report (individual CSV files or zipped files of
	// individual files).
	Format BusinessReportFormat
}

// The skill store category that is shown. Alexa skills are assigned a specific
// skill category during creation, such as News, Social, and Sports.
type Category struct {
	// The ID of the skill store category.
	CategoryId *int64
	// The name of the skill store category.
	CategoryName *string
}

// The default conference provider that is used if no other scheduled meetings are
// detected.
type ConferencePreference struct {
	// The ARN of the default conference provider.
	DefaultConferenceProviderArn *string
}

// An entity that provides a conferencing solution. Alexa for Business acts as the
// voice interface and mediator that connects users to their preferred conference
// provider. Examples of conference providers include Amazon Chime, Zoom, Cisco,
// and Polycom.
type ConferenceProvider struct {
	// The IP endpoint and protocol for calling.
	IPDialIn *IPDialIn
	// The type of conference providers.
	Type ConferenceProviderType
	// The name of the conference provider.
	Name *string
	// The meeting settings for the conference provider.
	MeetingSetting *MeetingSetting
	// The information for PSTN conferencing.
	PSTNDialIn *PSTNDialIn
	// The ARN of the newly created conference provider.
	Arn *string
}

// A contact with attributes.
type Contact struct {
	// The phone number of the contact. The phone number type defaults to WORK. You can
	// either specify PhoneNumber or PhoneNumbers. We recommend that you use
	// PhoneNumbers, which lets you specify the phone number type and multiple numbers.
	PhoneNumber *string
	// The list of SIP addresses for the contact.
	SipAddresses []*SipAddress
	// The name of the contact to display on the console.
	DisplayName *string
	// The ARN of the contact.
	ContactArn *string
	// The last name of the contact, used to call the contact on the device.
	LastName *string
	// The first name of the contact, used to call the contact on the device.
	FirstName *string
	// The list of phone numbers for the contact.
	PhoneNumbers []*PhoneNumber
}

// Information related to a contact.
type ContactData struct {
	// The phone number of the contact. The phone number type defaults to WORK. You can
	// specify PhoneNumber or PhoneNumbers. We recommend that you use PhoneNumbers,
	// which lets you specify the phone number type and multiple numbers.
	PhoneNumber *string
	// The name of the contact to display on the console.
	DisplayName *string
	// The last name of the contact, used to call the contact on the device.
	LastName *string
	// The list of SIP addresses for the contact.
	SipAddresses []*SipAddress
	// The ARN of the contact.
	ContactArn *string
	// The list of phone numbers for the contact.
	PhoneNumbers []*PhoneNumber
	// The first name of the contact, used to call the contact on the device.
	FirstName *string
}

// The content definition. This can contain only one text, SSML, or audio list
// object.
type Content struct {
	// The list of audio messages.
	AudioList []*Audio
	// The list of text messages.
	TextList []*Text
	// The list of SSML messages.
	SsmlList []*Ssml
}

// Creates settings for the end of meeting reminder feature that are applied to a
// room profile. The end of meeting reminder enables Alexa to remind users when a
// meeting is ending.
type CreateEndOfMeetingReminder struct {
	// The type of sound that users hear during the end of meeting reminder.
	ReminderType EndOfMeetingReminderType
	// A range of 3 to 15 minutes that determines when the reminder begins.
	ReminderAtMinutes []*int32
	// Whether an end of meeting reminder is enabled or not.
	Enabled *bool
}

// Creates settings for the instant booking feature that are applied to a room
// profile. When users start their meeting with Alexa, Alexa automatically books
// the room for the configured duration if the room is available.
type CreateInstantBooking struct {
	// Duration between 15 and 240 minutes at increments of 15 that determines how long
	// to book an available room when a meeting is started with Alexa.
	DurationInMinutes *int32
	// Whether instant booking is enabled or not.
	Enabled *bool
}

// Creates meeting room settings of a room profile.
type CreateMeetingRoomConfiguration struct {
	// Settings to automatically book a room for a configured duration if it's free
	// when joining a meeting with Alexa.
	InstantBooking *CreateInstantBooking
	// Settings for requiring a check in when a room is reserved. Alexa can cancel a
	// room reservation if it's not checked into to make the room available for others.
	// Users can check in by joining the meeting with Alexa or an AVS device, or by
	// saying “Alexa, check in.”
	RequireCheckIn *CreateRequireCheckIn
	// Whether room utilization metrics are enabled or not.
	RoomUtilizationMetricsEnabled *bool
	// Creates settings for the end of meeting reminder feature that are applied to a
	// room profile. The end of meeting reminder enables Alexa to remind users when a
	// meeting is ending.
	EndOfMeetingReminder *CreateEndOfMeetingReminder
}

// Creates settings for the require check in feature that are applied to a room
// profile. Require check in allows a meeting room’s Alexa or AVS device to prompt
// the user to check in; otherwise, the room will be released.
type CreateRequireCheckIn struct {
	// Duration between 5 and 20 minutes to determine when to release the room if it's
	// not checked into.
	ReleaseAfterMinutes *int32
	// Whether require check in is enabled or not.
	Enabled *bool
}

// The details about the developer that published the skill.
type DeveloperInfo struct {
	// The email of the developer.
	Email *string
	// The name of the developer.
	DeveloperName *string
	// The URL of the privacy policy.
	PrivacyPolicy *string
	// The website of the developer.
	Url *string
}

// A device with attributes.
type Device struct {
	// The ARN of a device.
	DeviceArn *string
	// The serial number of a device.
	DeviceSerialNumber *string
	// The status of a device. If the status is not READY, check the DeviceStatusInfo
	// value for details.
	DeviceStatus DeviceStatus
	// The name of a device.
	DeviceName *string
	// The MAC address of a device.
	MacAddress *string
	// Detailed information about a device's status.
	DeviceStatusInfo *DeviceStatusInfo
	// The room ARN of a device.
	RoomArn *string
	// The type of a device.
	DeviceType *string
	// The software version of a device.
	SoftwareVersion *string
	// Detailed information about a device's network profile.
	NetworkProfileInfo *DeviceNetworkProfileInfo
}

// Device attributes.
type DeviceData struct {
	// The name of the room associated with a device.
	RoomName *string
	// The room ARN associated with a device.
	RoomArn *string
	// Detailed information about a device's status.
	DeviceStatusInfo *DeviceStatusInfo
	// The name of a device.
	DeviceName *string
	// The status of a device.
	DeviceStatus DeviceStatus
	// The ARN of a device.
	DeviceArn *string
	// The name of the network profile associated with a device.
	NetworkProfileName *string
	// The type of a device.
	DeviceType *string
	// The software version of a device.
	SoftwareVersion *string
	// The time (in epoch) when the device data was created.
	CreatedTime *time.Time
	// The serial number of a device.
	DeviceSerialNumber *string
	// The MAC address of a device.
	MacAddress *string
	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string
}

// The list of device events.
type DeviceEvent struct {
	// The type of device event.
	Type DeviceEventType
	// The time (in epoch) when the event occurred.
	Timestamp *time.Time
	// The value of the event.
	Value *string
}

// Detailed information about a device's network profile.
type DeviceNetworkProfileInfo struct {
	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string
	// The time (in epoch) when the certificate expires.
	CertificateExpirationTime *time.Time
	// The ARN of the certificate associated with a device.
	CertificateArn *string
}

// Details of a device’s status.
type DeviceStatusDetail struct {
	// The list of available features on the device.
	Feature Feature
	// The device status detail code.
	Code DeviceStatusDetailCode
}

// Detailed information about a device's status.
type DeviceStatusInfo struct {
	// The time (in epoch) when the device connection status changed.
	ConnectionStatusUpdatedTime *time.Time
	// The latest available information about the connection status of a device.
	ConnectionStatus ConnectionStatus
	// One or more device status detail descriptions.
	DeviceStatusDetails []*DeviceStatusDetail
}

// Settings for the end of meeting reminder feature that are applied to a room
// profile. The end of meeting reminder enables Alexa to remind users when a
// meeting is ending.
type EndOfMeetingReminder struct {
	// Whether an end of meeting reminder is enabled or not.
	Enabled *bool
	// The type of sound that users hear during the end of meeting reminder.
	ReminderType EndOfMeetingReminderType
	// A range of 3 to 15 minutes that determines when the reminder begins.
	ReminderAtMinutes []*int32
}

// A filter name and value pair that is used to return a more specific list of
// results. Filters can be used to match a set of resources by various criteria.
type Filter struct {
	// The values of a filter.
	Values []*string
	// The key of a filter.
	Key *string
}

// The details of the gateway.
type Gateway struct {
	// The ARN of the gateway.
	Arn *string
	// The description of the gateway.
	Description *string
	// The ARN of the gateway group that the gateway is associated to.
	GatewayGroupArn *string
	// The software version of the gateway. The gateway automatically updates its
	// software version during normal operation.
	SoftwareVersion *string
	// The name of the gateway.
	Name *string
}

// The details of the gateway group.
type GatewayGroup struct {
	// The description of the gateway group.
	Description *string
	// The name of the gateway group.
	Name *string
	// The ARN of the gateway group.
	Arn *string
}

// The summary of a gateway group.
type GatewayGroupSummary struct {
	// The name of the gateway group.
	Name *string
	// The ARN of the gateway group.
	Arn *string
	// The description of the gateway group.
	Description *string
}

// The summary of a gateway.
type GatewaySummary struct {
	// The description of the gateway.
	Description *string
	// The ARN of the gateway.
	Arn *string
	// The ARN of the gateway group that the gateway is associated to.
	GatewayGroupArn *string
	// The name of the gateway.
	Name *string
	// The software version of the gateway. The gateway automatically updates its
	// software version during normal operation.
	SoftwareVersion *string
}

// Settings for the instant booking feature that are applied to a room profile.
// When users start their meeting with Alexa, Alexa automatically books the room
// for the configured duration if the room is available.
type InstantBooking struct {
	// Whether instant booking is enabled or not.
	Enabled *bool
	// Duration between 15 and 240 minutes at increments of 15 that determines how long
	// to book an available room when a meeting is started with Alexa.
	DurationInMinutes *int32
}

// The IP endpoint and protocol for calling.
type IPDialIn struct {
	// The IP address.
	Endpoint *string
	// The protocol, including SIP, SIPS, and H323.
	CommsProtocol CommsProtocol
}

// Meeting room settings of a room profile.
type MeetingRoomConfiguration struct {
	// Settings to automatically book the room if available for a configured duration
	// when joining a meeting with Alexa.
	InstantBooking *InstantBooking
	// Whether room utilization metrics are enabled or not.
	RoomUtilizationMetricsEnabled *bool
	// Settings for requiring a check in when a room is reserved. Alexa can cancel a
	// room reservation if it's not checked into. This makes the room available for
	// others. Users can check in by joining the meeting with Alexa or an AVS device,
	// or by saying “Alexa, check in.”
	RequireCheckIn *RequireCheckIn
	// Settings for the end of meeting reminder feature that are applied to a room
	// profile. The end of meeting reminder enables Alexa to remind users when a
	// meeting is ending.
	EndOfMeetingReminder *EndOfMeetingReminder
}

// The values that indicate whether a pin is always required (YES), never required
// (NO), or OPTIONAL.
//
//     * If YES, Alexa will always ask for a meeting pin.
//
//
// * If NO, Alexa will never ask for a meeting pin.
//
//     * If OPTIONAL, Alexa will
// ask if you have a meeting pin and if the customer responds with yes, it will ask
// for the meeting pin.
type MeetingSetting struct {
	// The values that indicate whether the pin is always required.
	RequirePin RequirePin
}

// The network profile associated with a device.
type NetworkProfile struct {
	// The root certificates of your authentication server, which is installed on your
	// devices and used to trust your authentication server during EAP negotiation.
	TrustAnchors []*string
	// The authentication standard that is used in the EAP framework. Currently,
	// EAP_TLS is supported.
	EapMethod NetworkEapMethod
	// The next, or subsequent, password of the Wi-Fi network. This password is
	// asynchronously transmitted to the device and is used when the password of the
	// network changes to NextPassword.
	NextPassword *string
	// Detailed information about a device's network profile.
	Description *string
	// The current password of the Wi-Fi network.
	CurrentPassword *string
	// The name of the network profile associated with a device.
	NetworkProfileName *string
	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string
	// The security type of the Wi-Fi network. This can be WPA2_ENTERPRISE, WPA2_PSK,
	// WPA_PSK, WEP, or OPEN.
	SecurityType NetworkSecurityType
	// The SSID of the Wi-Fi network.
	Ssid *string
	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string
}

// The data associated with a network profile.
type NetworkProfileData struct {
	// The SSID of the Wi-Fi network.
	Ssid *string
	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string
	// The authentication standard that is used in the EAP framework. Currently,
	// EAP_TLS is supported.
	EapMethod NetworkEapMethod
	// Detailed information about a device's network profile.
	Description *string
	// The security type of the Wi-Fi network. This can be WPA2_ENTERPRISE, WPA2_PSK,
	// WPA_PSK, WEP, or OPEN.
	SecurityType NetworkSecurityType
	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string
	// The name of the network profile associated with a device.
	NetworkProfileName *string
}

// The phone number for the contact containing the raw number and phone number
// type.
type PhoneNumber struct {
	// The raw value of the phone number.
	Number *string
	// The type of the phone number.
	Type PhoneNumberType
}

// A room profile with attributes.
type Profile struct {
	// The time zone of a room profile.
	Timezone *string
	// Retrieves if the profile is default or not.
	IsDefault *bool
	// The locale of a room profile. (This is currently available only to a limited
	// preview audience.)
	Locale *string
	// The wake word of a room profile.
	WakeWord WakeWord
	// The ARN of a room profile.
	ProfileArn *string
	// The temperature unit of a room profile.
	TemperatureUnit TemperatureUnit
	// The name of a room profile.
	ProfileName *string
	// The max volume limit of a room profile.
	MaxVolumeLimit *int32
	// Meeting room settings of a room profile.
	MeetingRoomConfiguration *MeetingRoomConfiguration
	// The ARN of the address book.
	AddressBookArn *string
	// The distance unit of a room profile.
	DistanceUnit DistanceUnit
	// The setup mode of a room profile.
	SetupModeDisabled *bool
	// The address of a room profile.
	Address *string
	// The PSTN setting of a room profile.
	PSTNEnabled *bool
}

// The data of a room profile.
type ProfileData struct {
	// The name of a room profile.
	ProfileName *string
	// The ARN of a room profile.
	ProfileArn *string
	// The locale of a room profile. (This is currently available only to a limited
	// preview audience.)
	Locale *string
	// The time zone of a room profile.
	Timezone *string
	// Retrieves if the profile data is default or not.
	IsDefault *bool
	// The address of a room profile.
	Address *string
	// The temperature unit of a room profile.
	TemperatureUnit TemperatureUnit
	// The distance unit of a room profile.
	DistanceUnit DistanceUnit
	// The wake word of a room profile.
	WakeWord WakeWord
}

// The information for public switched telephone network (PSTN) conferencing.
type PSTNDialIn struct {
	// The delay duration before Alexa enters the conference pin with dual-tone
	// multi-frequency (DTMF). Each number on the dial pad corresponds to a DTMF tone,
	// which is how we send data over the telephone network.
	OneClickPinDelay *string
	// The phone number to call to join the conference.
	PhoneNumber *string
	// The zip code.
	CountryCode *string
	// The delay duration before Alexa enters the conference ID with dual-tone
	// multi-frequency (DTMF). Each number on the dial pad corresponds to a DTMF tone,
	// which is how we send data over the telephone network.
	OneClickIdDelay *string
}

// Settings for the require check in feature that are applied to a room profile.
// Require check in allows a meeting room’s Alexa or AVS device to prompt the user
// to check in; otherwise, the room will be released.
type RequireCheckIn struct {
	// Whether require check in is enabled or not.
	Enabled *bool
	// Duration between 5 and 20 minutes to determine when to release the room if it's
	// not checked into.
	ReleaseAfterMinutes *int32
}

// A room with attributes.
type Room struct {
	// The ARN of a room.
	RoomArn *string
	// The provider calendar ARN of a room.
	ProviderCalendarId *string
	// The description of a room.
	Description *string
	// The profile ARN of a room.
	ProfileArn *string
	// The name of a room.
	RoomName *string
}

// The data of a room.
type RoomData struct {
	// The profile name of a room.
	ProfileName *string
	// The name of a room.
	RoomName *string
	// The ARN of a room.
	RoomArn *string
	// The provider calendar ARN of a room.
	ProviderCalendarId *string
	// The description of a room.
	Description *string
	// The profile ARN of a room.
	ProfileArn *string
}

// A skill parameter associated with a room.
type RoomSkillParameter struct {
	// The parameter key of a room skill parameter. ParameterKey is an enumerated type
	// that only takes “DEFAULT” or “SCOPE” as valid values.
	ParameterKey *string
	// The parameter value of a room skill parameter.
	ParameterValue *string
}

// The SIP address for the contact containing the URI and SIP address type.
type SipAddress struct {
	// The URI for the SIP address.
	Uri *string
	// The type of the SIP address.
	Type SipType
}

// Granular information about the skill.
type SkillDetails struct {
	// The details about what the skill supports organized as bullet points.
	BulletPoints []*string
	// The types of skills.
	SkillTypes []*string
	// The updates added in bullet points.
	NewInThisVersionBulletPoints []*string
	// The date when the skill was released.
	ReleaseDate *string
	// The details about the developer that published the skill.
	DeveloperInfo *DeveloperInfo
	// The description of the product.
	ProductDescription *string
	// This member has been deprecated. The list of reviews for the skill, including
	// Key and Value pair.
	Reviews map[string]*string
	// The phrase used to trigger the skill.
	InvocationPhrase *string
	// The generic keywords associated with the skill that can be used to find a skill.
	GenericKeywords []*string
	// The URL of the end user license agreement.
	EndUserLicenseAgreement *string
}

// A skill group with attributes.
type SkillGroup struct {
	// The name of a skill group.
	SkillGroupName *string
	// The ARN of a skill group.
	SkillGroupArn *string
	// The description of a skill group.
	Description *string
}

// The attributes of a skill group.
type SkillGroupData struct {
	// The description of a skill group.
	Description *string
	// The skill group ARN of a skill group.
	SkillGroupArn *string
	// The skill group name of a skill group.
	SkillGroupName *string
}

// The detailed information about an Alexa skill.
type SkillsStoreSkill struct {
	// The URL where the skill icon resides.
	IconUrl *string
	// The name of the skill.
	SkillName *string
	// Sample utterances that interact with the skill.
	SampleUtterances []*string
	// Linking support for a skill.
	SupportsLinking *bool
	// The ARN of the skill.
	SkillId *string
	// Short description about the skill.
	ShortDescription *string
	// Information about the skill.
	SkillDetails *SkillDetails
}

// The summary of skills.
type SkillSummary struct {
	// The ARN of the skill summary.
	SkillId *string
	// Whether the skill is enabled under the user's account, or if it requires linking
	// to be used.
	EnablementType EnablementType
	// Linking support for a skill.
	SupportsLinking *bool
	// Whether the skill is publicly available or is a private skill.
	SkillType SkillType
	// The name of the skill.
	SkillName *string
}

// A smart home appliance that can connect to a central system. Any domestic device
// can be a smart appliance.
type SmartHomeAppliance struct {
	// The friendly name of the smart home appliance.
	FriendlyName *string
	// The description of the smart home appliance.
	Description *string
	// The name of the manufacturer of the smart home appliance.
	ManufacturerName *string
}

// An object representing a sort criteria.
type Sort struct {
	// The sort key of a sort object.
	Key *string
	// The sort value of a sort object.
	Value SortValue
}

// The SSML message. For more information, see SSML Reference
// (https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html).
type Ssml struct {
	// The value of the SSML message in the correct SSML format. The audio tag is not
	// supported.
	Value *string
	// The locale of the SSML message. Currently, en-US is supported.
	Locale Locale
}

// A key-value pair that can be associated with a resource.
type Tag struct {
	// The value of a tag. Tag values are case sensitive and can be null.
	Value *string
	// The key of a tag. Tag keys are case-sensitive.
	Key *string
}

// The text message.
type Text struct {
	// The value of the text message.
	Value *string
	// The locale of the text message. Currently, en-US is supported.
	Locale Locale
}

// Settings for the end of meeting reminder feature that are applied to a room
// profile. The end of meeting reminder enables Alexa to remind users when a
// meeting is ending.
type UpdateEndOfMeetingReminder struct {
	// Updates settings for the end of meeting reminder feature that are applied to a
	// room profile. The end of meeting reminder enables Alexa to remind users when a
	// meeting is ending.
	ReminderAtMinutes []*int32
	// The type of sound that users hear during the end of meeting reminder.
	ReminderType EndOfMeetingReminderType
	// Whether an end of meeting reminder is enabled or not.
	Enabled *bool
}

// Updates settings for the instant booking feature that are applied to a room
// profile. If instant booking is enabled, Alexa automatically reserves a room if
// it is free when a user joins a meeting with Alexa.
type UpdateInstantBooking struct {
	// Whether instant booking is enabled or not.
	Enabled *bool
	// Duration between 15 and 240 minutes at increments of 15 that determines how long
	// to book an available room when a meeting is started with Alexa.
	DurationInMinutes *int32
}

// Updates meeting room settings of a room profile.
type UpdateMeetingRoomConfiguration struct {
	// Settings for the end of meeting reminder feature that are applied to a room
	// profile. The end of meeting reminder enables Alexa to remind users when a
	// meeting is ending.
	EndOfMeetingReminder *UpdateEndOfMeetingReminder
	// Settings for requiring a check in when a room is reserved. Alexa can cancel a
	// room reservation if it's not checked into to make the room available for others.
	// Users can check in by joining the meeting with Alexa or an AVS device, or by
	// saying “Alexa, check in.”
	RequireCheckIn *UpdateRequireCheckIn
	// Whether room utilization metrics are enabled or not.
	RoomUtilizationMetricsEnabled *bool
	// Settings to automatically book an available room available for a configured
	// duration when joining a meeting with Alexa.
	InstantBooking *UpdateInstantBooking
}

// Updates settings for the require check in feature that are applied to a room
// profile. Require check in allows a meeting room’s Alexa or AVS device to prompt
// the user to check in; otherwise, the room will be released.
type UpdateRequireCheckIn struct {
	// Duration between 5 and 20 minutes to determine when to release the room if it's
	// not checked into.
	ReleaseAfterMinutes *int32
	// Whether require check in is enabled or not.
	Enabled *bool
}

// Information related to a user.
type UserData struct {
	// The email of a user.
	Email *string
	// The enrollment ARN of a user.
	EnrollmentId *string
	// The first name of a user.
	FirstName *string
	// The last name of a user.
	LastName *string
	// The enrollment status of a user.
	EnrollmentStatus EnrollmentStatus
	// The ARN of a user.
	UserArn *string
}
