import styled from 'styled-components';
import {
  colorText,
  colorGrayLighter,
  colorDanger,
  colorGrayDark,
  colorGrayLightest,
} from '/imports/ui/stylesheets/styled-components/palette';
import {
  smPaddingX,
  xsPadding,
  smPaddingY,
  borderRadius,
  borderSize,
} from '/imports/ui/stylesheets/styled-components/general';
import { fontSizeBase } from '/imports/ui/stylesheets/styled-components/typography';
import TextareaAutosize from 'react-autosize-textarea';
import EmojiPickerComponent from '/imports/ui/components/emoji-picker/component';
import Button from '@mui/material/Button';

interface FormProps {
  isRTL: boolean;
}

const Form = styled.form<FormProps>`
  flex-grow: 0;
  flex-shrink: 0;
  align-self: flex-end;
  width: 100%;
  position: relative;
  margin-top: .2rem;
`;

const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  border-radius: 0.75rem;
`;

const Input = styled(TextareaAutosize)`
  flex: 1;
  background: #fff;
  background-clip: padding-box;
  margin: ${xsPadding} 0 ${xsPadding} ${xsPadding};
  color: ${colorGrayLighter};
  -webkit-appearance: none;
  padding: calc(${smPaddingY} * 2.5) 0 calc(${smPaddingX} * 1.25) calc(${smPaddingY} * 2.5);
  resize: none;
  transition: color 0.3s ease;
  border-radius: ${borderRadius};
  font-size: ${fontSizeBase};
  line-height: 1;
  min-height: 2.5rem;
  max-height: 3.5rem;
  overflow-y: auto;
  border: ${colorGrayLightest};
  box-shadow: none;
  outline: none;

  [dir='ltr'] & {
    border-radius: 0.75rem 0 0 0.75rem;
  }

  [dir='rtl'] & {
    border-radius: 0 0.75rem 0.75rem 0;
  }

  &:focus {
    color: ${colorText};
  }

  &:disabled,
  &[disabled] {
    cursor: not-allowed;
    opacity: .75;
    background-color: rgba(167,179,189,0.25);
  }
`;

// @ts-ignore - as button comes from JS, we can't provide its props
const SendButton = styled(Button)`
  align-self: center;
  font-size: 0.9rem;
  height: 100%;

  & > span {
    height: 100%;
    display: flex;
    align-items: center;
    border-radius: 0 0.75rem 0.75rem 0;
  }

  [dir="rtl"]  & {
    -webkit-transform: scale(-1, 1);
    -moz-transform: scale(-1, 1);
    -ms-transform: scale(-1, 1);
    -o-transform: scale(-1, 1);
    transform: scale(-1, 1);
  }
`;

const EmojiButtonWrapper = styled.div``;

// @ts-ignore - as button comes from JS, we can't provide its props
const EmojiButton = styled(Button)`
  margin:0 0 0 ${smPaddingX};

  [dir="rtl"]  & {
    margin: 0 ${smPaddingX} 0 0;
    -webkit-transform: scale(-1, 1);
    -moz-transform: scale(-1, 1);
    -ms-transform: scale(-1, 1);
    -o-transform: scale(-1, 1);
    transform: scale(-1, 1);
  }
`;

const EmojiPickerWrapper = styled.div`
position: absolute;
bottom: calc(100% + 0.5rem);
left: 0;
right: 0;
border: 1px solid ${colorGrayLighter};
border-radius: ${borderRadius};
box-shadow: 0 2px 10px rgba(0,0,0,0.1);
z-index: 1000;

  .emoji-mart {
    max-width: 100% !important;
  }
  .emoji-mart-anchor {
    cursor: pointer;
  }
  .emoji-mart-emoji {
    cursor: pointer !important;
  }
  .emoji-mart-category-list {
    span {
      cursor: pointer !important;
      display: inline-block !important;
    }
  }
`;

const ChatMessageError = styled.div`
  color: ${colorDanger};
  font-size: calc(${fontSizeBase} * .75);
  color: ${colorGrayDark};
  text-align: left;
  padding: ${borderSize} 0;
  word-break: break-word;
  position: relative;
  margin-right: 0.05rem;
  margin-left: 0.05rem;
`;

const EmojiPicker = styled(EmojiPickerComponent)`
  position: relative;
`;

const InputWrapper = styled.div`
  display: flex;
  flex-direction: row;
  flex-grow: 1;
  min-width: 0;
  z-index: 0;
  border-radius: 0.75rem;
  border: 1px solid ${colorGrayLighter};

  &:focus-within {
    box-shadow: 0 0 0 ${xsPadding} ${colorGrayLighter};
  }
`;

export default {
  Form,
  Wrapper,
  Input,
  SendButton,
  EmojiButton,
  EmojiButtonWrapper,
  EmojiPicker,
  EmojiPickerWrapper,
  ChatMessageError,
  InputWrapper,
};
