package org.bigbluebutton.core.apps.users

import org.bigbluebutton.common2.msgs._
import org.bigbluebutton.core.apps.RightsManagementTrait
import org.bigbluebutton.core.db.{ CaptionLocaleDAO, CaptionTypes }
import org.bigbluebutton.core.models.{ UserState, Users2x }
import org.bigbluebutton.core.running.{ LiveMeeting, OutMsgRouter }

trait SetUserSpeechLocaleMsgHdlr extends RightsManagementTrait {
  this: UsersApp =>

  val liveMeeting: LiveMeeting
  val outGW: OutMsgRouter

  def handleSetUserSpeechLocaleReqMsg(msg: SetUserSpeechLocaleReqMsg): Unit = {
    log.info("handleSetUserSpeechLocaleReqMsg: locale={} provider={} userId={}", msg.body.locale, msg.body.provider, msg.header.userId)

    def broadcastUserSpeechLocaleChanged(user: UserState, locale: String, provider: String): Unit = {
      val routingChange = Routing.addMsgToClientRouting(
        MessageTypes.BROADCAST_TO_MEETING,
        liveMeeting.props.meetingProp.intId, user.intId
      )
      val envelopeChange = BbbCoreEnvelope(UserSpeechLocaleChangedEvtMsg.NAME, routingChange)
      val headerChange = BbbClientMsgHeader(UserSpeechLocaleChangedEvtMsg.NAME, liveMeeting.props.meetingProp.intId, user.intId)

      val bodyChange = UserSpeechLocaleChangedEvtMsgBody(locale, provider)
      val eventChange = UserSpeechLocaleChangedEvtMsg(headerChange, bodyChange)
      val msgEventChange = BbbCommonEnvCoreMsg(envelopeChange, eventChange)
      outGW.send(msgEventChange)
    }

    for {
      user <- Users2x.findWithIntId(liveMeeting.users2x, msg.header.userId)
    } yield {
      var changeLocale: Option[UserState] = None;
      changeLocale = Users2x.setUserSpeechLocale(liveMeeting.users2x, msg.header.userId, msg.body.locale)

      // Add new CaptionLocale
      CaptionLocaleDAO.insertOrUpdateCaptionLocale(
        liveMeeting.props.meetingProp.intId,
        msg.body.locale, CaptionTypes.AUDIO_TRANSCRIPTION, msg.header.userId
      )
      broadcastUserSpeechLocaleChanged(user, msg.body.locale, msg.body.provider)
    }

  }
}
