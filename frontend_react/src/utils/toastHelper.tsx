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

export const showSuccessToast = (message: string) => {
    toast.success(message, {
        duration: defaultDuration,
    });
};

export const showErrorToast = (message: string) => {
    toast.error(message, {
        duration: defaultDuration,
    });
};

export const showWarningToast = (message: string) => {
    toast(message, {
        icon: "⚠️",
        duration: defaultDuration,
    });
};

export const showInfoToast = (message: string) => {
    toast(message, {
        icon: "ℹ️",
        duration: defaultDuration,
    });
};