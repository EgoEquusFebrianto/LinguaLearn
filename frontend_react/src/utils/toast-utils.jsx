import toast from 'react-hot-toast'

const defaultDuration = 2500;

export const showLoginRequiredToast  = () => {
  toast.error(
    "Please login to use this feature.",
    {
        duration: defaultDuration,
    }
  );
};

export const showSuccessToast = (message) => {
    toast.success(message, {
        duration: defaultDuration,
    });
};

export const showErrorToast = (message) => {
    toast.error(message, {
        duration: defaultDuration,
    });
};

export const showWarningToast = (message) => {
    toast(message, {
        icon: "⚠️",
        duration: defaultDuration,
    });
};

export const showInfoToast = (message) => {
    toast(message, {
        icon: "ℹ️",
        duration: defaultDuration,
    });
};